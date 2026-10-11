package com.ciagotech.kan.server.auth;

import org.junit.jupiter.api.Test;
import org.springframework.mock.web.MockHttpServletRequest;
import org.springframework.mock.web.MockHttpServletResponse;
import java.nio.file.Path;
import java.nio.charset.StandardCharsets;
import java.time.Instant;
import java.time.Clock;
import java.time.ZoneOffset;
import java.time.format.DateTimeFormatter;
import java.util.Map;
import java.util.concurrent.atomic.AtomicBoolean;
import static org.junit.jupiter.api.Assertions.*;

class AgentAuthSecurityTest {
    private final SigV1Rules rules = SigV1Rules.load(Path.of("../contracts/sigv1-rules.json"));
    AgentAuthSecurityTest() throws Exception {}
    private static final String KEY = "test-agent-key", SECRET = "test-secret";
    private KeyService keys(String kind) {
        return id -> KEY.equals(id) ? new KeyService.Credential("account-1",KEY,kind,SECRET) : null;
    }
    private MockHttpServletRequest signed(String nonce, byte[] body) {
        var request = new MockHttpServletRequest("PUT", "/internal/v1/agent/work");
        String date = DateTimeFormatter.ofPattern("uuuuMMdd'T'HHmmss'Z'").withZone(ZoneOffset.UTC).format(Instant.now());
        request.addHeader("Host", "localhost:4700");
        request.addHeader("X-Kan-Date", date);
        request.addHeader("X-Kan-Access-Key", KEY);
        request.addHeader("X-Kan-Nonce", nonce);
        request.addHeader("Content-Type", "application/json");
        request.setContent(body);
        request.addHeader("X-Kan-Signature", KanSigV1.sign(SECRET,"PUT",request.getRequestURI(),"",
            Map.of("host","localhost:4700","x-kan-date",date,"x-kan-access-key",KEY,
                   "x-kan-nonce",nonce,"content-type","application/json"),body,date,rules));
        return request;
    }
    private int run(KanSigV1AuthFilter filter, MockHttpServletRequest request) throws Exception {
        var response = new MockHttpServletResponse();
        filter.doFilter(request,response,(req,res) -> {
            assertInstanceOf(KeyService.Principal.class,req.getAttribute("kan.principal"));
            assertEquals("account-1",((KeyService.Principal)req.getAttribute("kan.principal")).accountId());
            res.getWriter().write("accepted");
        });
        return response.getStatus();
    }
    @Test void validReplayAndNonceTampering() throws Exception {
        var filter = new KanSigV1AuthFilter(rules,keys("service"),new InMemoryNonceStore());
        assertEquals(200,run(filter,signed("nonce-valid-00000001",new byte[0])));
        assertEquals(401,run(filter,signed("nonce-valid-00000001",new byte[0])));
        var changed = signed("nonce-valid-00000002",new byte[0]);
        changed.removeHeader("X-Kan-Nonce");changed.addHeader("X-Kan-Nonce","nonce-valid-00000003");
        assertEquals(401,run(filter,changed));
        assertEquals(200,run(filter,signed("nonce-valid-00000003",new byte[0])));
    }
    @Test void badSignatureCannotPoisonNonce() throws Exception {
        var filter = new KanSigV1AuthFilter(rules,keys("service"),new InMemoryNonceStore());
        var bad = signed("nonce-poison-0000001",new byte[0]);
        bad.removeHeader("X-Kan-Signature");bad.addHeader("X-Kan-Signature","0".repeat(64));
        assertEquals(401,run(filter,bad));
        assertEquals(200,run(filter,signed("nonce-poison-0000001",new byte[0])));
    }
    @Test void wrongKindAndMissingKeyRejected() throws Exception {
        assertEquals(401,run(new KanSigV1AuthFilter(rules,keys("user"),new InMemoryNonceStore()),signed("nonce-kind-000000001",new byte[0])));
        assertEquals(401,run(new KanSigV1AuthFilter(rules,id -> null,new InMemoryNonceStore()),signed("nonce-kind-000000002",new byte[0])));
    }
    @Test void bodyBoundIncludesUnknownLength() throws Exception {
        var request = signed("nonce-large-00000001",new byte[CachedBodyHttpServletRequest.MAX_BODY_BYTES+1]);
        var unknownLength = new jakarta.servlet.http.HttpServletRequestWrapper(request) {
            @Override public long getContentLengthLong() { return -1; }
            @Override public int getContentLength() { return -1; }
        };
        var response = new MockHttpServletResponse();
        new KanSigV1AuthFilter(rules,keys("service"),new InMemoryNonceStore()).doFilter(unknownLength,response,
            (req,res) -> fail("oversized body forwarded"));
        assertEquals(413,response.getStatus());
    }
    @Test void downstreamCanReadBodyAgain() throws Exception {
        var request = signed("nonce-body-000000001","{}".getBytes(StandardCharsets.UTF_8));
        var response = new MockHttpServletResponse();
        new KanSigV1AuthFilter(rules,keys("service"),new InMemoryNonceStore()).doFilter(request,response,
            (req,res) -> assertArrayEquals("{}".getBytes(StandardCharsets.UTF_8),req.getInputStream().readAllBytes()));
        assertEquals(200,response.getStatus());
    }
    @Test void duplicateHeadersAndStaleDateRejected() throws Exception {
        var filter = new KanSigV1AuthFilter(rules,keys("service"),new InMemoryNonceStore());
        var duplicate = signed("nonce-duplicate-0001",new byte[0]);duplicate.addHeader("X-Kan-Nonce","nonce-duplicate-0002");
        assertEquals(401,run(filter,duplicate));
        var stale = signed("nonce-stale-00000001",new byte[0]);stale.removeHeader("X-Kan-Date");stale.addHeader("X-Kan-Date","20260926T120000Z");
        assertEquals(401,run(filter,stale));
    }
    @Test void unavailableNonceStoreFailsClosed() throws Exception {
        NonceStore unavailable = (key,nonce,ttl) -> { throw new IllegalStateException("offline"); };
        assertEquals(503,run(new KanSigV1AuthFilter(rules,keys("service"),unavailable),signed("nonce-offline-000001",new byte[0])));
    }
    @Test void expiringStoreAndCapacity() {
        Instant now = Instant.parse("2026-10-11T00:00:00Z");
        var clock = new Clock() {
            private int calls;
            @Override public java.time.ZoneId getZone() { return ZoneOffset.UTC; }
            @Override public Clock withZone(java.time.ZoneId zone) { return this; }
            @Override public Instant instant() { return now.plusSeconds(calls++ < 2 ? 0 : 3); }
        };
        var store = new InMemoryNonceStore(clock,1);
        assertTrue(store.checkAndStore("key","nonce",2));
        assertFalse(store.checkAndStore("key","nonce",2));
        assertTrue(store.checkAndStore("key","nonce",2));
        assertThrows(IllegalStateException.class,() -> store.checkAndStore("key","other",2));
    }
    @Test void missingCanonicalHeaderRejected() {
        assertThrows(IllegalArgumentException.class,() -> KanSigV1.buildCanonicalRequest("GET","/","",Map.of(),new byte[0],rules));
    }
}

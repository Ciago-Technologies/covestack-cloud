package com.ciagotech.kan.server.auth;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.core.StringRedisTemplate;
import java.io.IOException;
import java.nio.file.Path;

/** Requires Person C's KeyService bean; no insecure fallback credentials. */
@Configuration
public class AgentAuthConfiguration {
    @Bean
    SigV1Rules sigV1Rules(@Value("${kan.sigv1.rules-path:../contracts/sigv1-rules.json}") String path)
            throws IOException {
        SigV1Rules rules = SigV1Rules.load(Path.of(path));
        if (!"KAN1-HMAC-SHA256".equals(rules.algorithm())
                || !"X-Kan-Date".equals(rules.dateHeader())
                || !"X-Kan-Access-Key".equals(rules.accessKeyHeader())
                || !"X-Kan-Nonce".equals(rules.nonceHeader())
                || !"X-Kan-Signature".equals(rules.signatureHeader())
                || !rules.signBodyContentTypeHeader()
                || !rules.alwaysSignedHeaders().containsAll(java.util.List.of(
                "host", "x-kan-date", "x-kan-access-key", "x-kan-nonce"))
                || rules.maxClockSkewSeconds() < 1
                || rules.nonceReplayWindowSeconds() < 2L * rules.maxClockSkewSeconds())
            throw new IllegalArgumentException("Unsafe signing rules");
        return rules;
    }
    @Bean
    NonceStore nonceStore(StringRedisTemplate redis) { return new ValkeyNonceStore(redis); }
    @Bean
    FilterRegistrationBean<KanSigV1AuthFilter> agentAuthFilter(
            SigV1Rules rules, KeyService keys, NonceStore nonces) {
        var registration = new FilterRegistrationBean<>(new KanSigV1AuthFilter(rules, keys, nonces));
        registration.setName("kanAgentAuth");
        registration.addUrlPatterns("/internal/v1/agent", "/internal/v1/agent/*");
        registration.setOrder(-100);
        return registration;
    }
    @Bean
    FilterRegistrationBean<BearerAuthFilter> publicAuthFilter(BearerTokenService tokens) {
        var registration = new FilterRegistrationBean<>(new BearerAuthFilter(tokens));
        registration.setName("kanPublicAuth");
        registration.addUrlPatterns("/v2", "/v2/*", "/kan/v1", "/kan/v1/*");
        registration.setOrder(-100);
        return registration;
    }

}

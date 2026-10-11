package com.ciagotech.kan.server.auth;

import org.junit.jupiter.api.Test;
import org.springframework.mock.web.MockHttpServletRequest;
import org.springframework.mock.web.MockHttpServletResponse;
import static org.junit.jupiter.api.Assertions.*;

class BearerAuthFilterTest {
    private static final String TOKEN = "example-opaque-token-0000000000000001";
    private int run(BearerTokenService tokens,String... values) throws Exception {
        var request=new MockHttpServletRequest("GET","/v2/droplets");
        for(String value:values) request.addHeader("Authorization",value);
        var response=new MockHttpServletResponse();
        new BearerAuthFilter(tokens).doFilter(request,response,(req,res) -> {
            assertEquals("account-1",((KeyService.Principal)req.getAttribute("kan.principal")).accountId());
        });
        return response.getStatus();
    }
    @Test void missingUnknownAndDuplicateRejected() throws Exception {
        assertEquals(401,run(token -> null));
        assertEquals(401,run(token -> null,"Bearer "+TOKEN));
        assertEquals(401,run(token -> new KeyService.Principal("account-1","user-key","user"),"Bearer "+TOKEN,"Bearer "+TOKEN));
    }
    @Test void activeUserAcceptedServiceRejected() throws Exception {
        assertEquals(200,run(token -> TOKEN.equals(token) ? new KeyService.Principal("account-1","user-key","user") : null,"Bearer "+TOKEN));
        assertEquals(401,run(token -> new KeyService.Principal("account-1","agent-key","service"),"Bearer "+TOKEN));
    }
}

package com.ciagotech.kan.server.auth;

import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.web.filter.OncePerRequestFilter;
import java.io.IOException;
import java.util.Enumeration;

/** Public API authentication; controllers separately enforce account/resource permissions. */
public final class BearerAuthFilter extends OncePerRequestFilter {
    private final BearerTokenService tokens;
    public BearerAuthFilter(BearerTokenService tokens) { this.tokens = tokens; }
    @Override protected void doFilterInternal(HttpServletRequest request,HttpServletResponse response,
            FilterChain chain) throws IOException,ServletException {
        Enumeration<String> values = request.getHeaders("Authorization");
        String value = values != null && values.hasMoreElements() ? values.nextElement() : null;
        if (value == null || values.hasMoreElements() || value.length() > 1024
                || !value.regionMatches(true,0,"Bearer ",0,7)) {
            reject(response);return;
        }
        String token = value.substring(7);
        if (!token.matches("[A-Za-z0-9._~+/-]{32,512}={0,2}")) { reject(response);return; }
        KeyService.Principal principal = tokens.findActivePrincipal(token);
        if (principal == null || principal.accountId() == null || principal.accountId().isBlank()
                || !"user".equals(principal.kind())) { reject(response);return; }
        request.setAttribute("kan.principal",principal);
        chain.doFilter(request,response);
    }
    private static void reject(HttpServletResponse response) throws IOException {
        response.setStatus(401);response.setHeader("WWW-Authenticate","Bearer");
        response.setContentType("application/json");response.getWriter().write("{\"error\":\"invalid bearer credential\"}");
    }
}

package com.ciagotech.kan.server.auth;

import jakarta.servlet.ReadListener;
import jakarta.servlet.ServletInputStream;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletRequestWrapper;

import java.io.BufferedReader;
import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.nio.charset.StandardCharsets;

/**
 * Reads the body ONCE here (to hash it for signature verification) and
 * caches it, so Spring's Jackson message converter can still read a full
 * body afterward. Without this, the filter's read empties the stream and
 * every controller behind it sees an empty request body. This is the
 * "cached body!" note your own TDD already put next to C3 — this class
 * is that fix.
 */
public class CachedBodyHttpServletRequest extends HttpServletRequestWrapper {

    public static final int MAX_BODY_BYTES = 1024 * 1024;

    public static class BodyTooLargeException extends IOException {}

    private final byte[] cachedBody;

    public CachedBodyHttpServletRequest(HttpServletRequest request) throws IOException {
        super(request);
        if (request.getContentLengthLong() > MAX_BODY_BYTES) throw new BodyTooLargeException();
        this.cachedBody = request.getInputStream().readNBytes(MAX_BODY_BYTES + 1);
        if (cachedBody.length > MAX_BODY_BYTES) throw new BodyTooLargeException();
    }

    public byte[] getCachedBody() {
        return cachedBody;
    }

    @Override
    public ServletInputStream getInputStream() {
        ByteArrayInputStream byteStream = new ByteArrayInputStream(cachedBody);
        return new ServletInputStream() {
            @Override public boolean isFinished() { return byteStream.available() == 0; }
            @Override public boolean isReady() { return true; }
            @Override public void setReadListener(ReadListener listener) {
                throw new UnsupportedOperationException("Synchronous body wrapper");
            }
            @Override public int read() { return byteStream.read(); }
        };
    }

    @Override
    public BufferedReader getReader() {
        return new BufferedReader(new InputStreamReader(getInputStream(), StandardCharsets.UTF_8));
    }
}

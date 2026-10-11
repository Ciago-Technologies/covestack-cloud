package com.ciagotech.kan.server.auth;

import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;

/**
 * Loaded from contracts/sigv1-rules.json — the same file Go reads.
 * Change header names or timing windows in ONE place, not two.
 */
public record SigV1Rules(
        String algorithm,
        String dateHeader,
        String accessKeyHeader,
        String signatureHeader,
        String nonceHeader,
        List<String> alwaysSignedHeaders,
        boolean signBodyContentTypeHeader,
        int maxClockSkewSeconds,
        int nonceReplayWindowSeconds,
        String dateFormat
) {
    public static SigV1Rules load(Path path) throws IOException {
        ObjectMapper mapper = new ObjectMapper();
        return mapper.readValue(Files.readString(path), SigV1Rules.class);
    }
}

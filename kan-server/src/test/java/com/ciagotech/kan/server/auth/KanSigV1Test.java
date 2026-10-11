package com.ciagotech.kan.server.auth;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;

class KanSigV1Test {

    // Adjust these two paths if your repo layout puts kan-server/ somewhere
    // other than next to contracts/.
    private static final Path RULES_PATH = Path.of("../contracts/sigv1-rules.json");
    private static final Path VECTORS_PATH = Path.of("../contracts/sigv1-vectors.json");

    @Test
    void vectorsMatch() throws Exception {
        SigV1Rules rules = SigV1Rules.load(RULES_PATH);
        ObjectMapper mapper = new ObjectMapper();
        JsonNode root = mapper.readTree(Files.readString(VECTORS_PATH));

        for (JsonNode v : root.get("vectors")) {
            Map<String, String> headers = new HashMap<>();
            Iterator<Map.Entry<String, JsonNode>> fields = v.get("headers").fields();
            while (fields.hasNext()) {
                var e = fields.next();
                headers.put(e.getKey(), e.getValue().asText());
            }
            String dateStr = headers.get("X-Kan-Date");
            byte[] body = v.get("body").asText().getBytes();
            String secret = v.get("secretKey").asText();
            String expected = v.get("expectedSignature").asText();

            String actual = KanSigV1.sign(secret, v.get("method").asText(), v.get("path").asText(),
                    v.get("query").asText(), headers, body, dateStr, rules);

            assertEquals(expected, actual, "mismatch for vector: " + v.get("name").asText());
        }
    }
}

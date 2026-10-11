package com.ciagotech.kan.server.contract;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;
import java.nio.file.Path;
import static org.junit.jupiter.api.Assertions.assertEquals;

class LifecycleValidatorTest {
    @Test void sharedLifecycleVectors() throws Exception {
        JsonNode root = new ObjectMapper().readTree(Path.of("../contracts/lifecycle-test-vectors.json").toFile());
        for (JsonNode v : root.path("vectors")) {
            String kind=v.path("kind").asText(), desired=v.path("desiredState").asText();
            String previous=v.path("previousState").asText(), state=v.path("state").asText();
            String result="invalid";
            try {
                switch (v.path("type").asText()) {
                    case "resource" -> {
                        LifecycleValidator.validateResource(kind,desired,v.path("deleting").asBoolean(),v.path("agentWork").asBoolean());result="valid";
                    }
                    case "observation" -> {
                        Long observed=v.path("observedGeneration").isNull() || v.path("observedGeneration").isMissingNode() ? null : v.path("observedGeneration").asLong();
                        result=LifecycleValidator.validateObservation(kind,desired,previous,state,v.path("targetGeneration").asLong(),observed) ? "applied" : "ignored";
                    }
                    case "server-observation" -> {
                        Long observed=v.path("observedGeneration").isNull() || v.path("observedGeneration").isMissingNode() ? null : v.path("observedGeneration").asLong();
                        result=LifecycleValidator.validateServerObservation(kind,desired,previous,state,v.path("targetGeneration").asLong(),observed) ? "applied" : "ignored";
                    }
                    case "transition" -> result=LifecycleValidator.canTransition(kind,previous,state) ? "allowed" : "denied";
                    case "entity" -> result=LifecycleValidator.canEntityTransition(v.path("entity").asText(),previous,state) ? "allowed" : "denied";
                    default -> throw new AssertionError("unknown vector type");
                }
            } catch (IllegalArgumentException ignored) { result="invalid"; }
            assertEquals(v.path("result").asText(),result,v.path("name").asText());
        }
    }
}

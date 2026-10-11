package com.ciagotech.kan.server.contract;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.util.ArrayDeque;
import java.util.HashSet;

/** Shared state validation; services additionally enforce account, lease, readiness and timestamp ordering. */
public final class LifecycleValidator {
    private LifecycleValidator() {}
    private static final JsonNode CATALOG = load();
    private static JsonNode load() {
        try { return new ObjectMapper().readTree(LifecycleData.JSON); }
        catch (Exception e) { throw new ExceptionInInitializerError(e); }
    }
    private static boolean contains(JsonNode values, String target) {
        if (values == null || !values.isArray()) return false;
        for (JsonNode value : values) if (value.asText().equals(target)) return true;
        return false;
    }
    private static JsonNode resource(String kind) {
        if (kind == null) throw new IllegalArgumentException("missing kind");
        JsonNode rule = CATALOG.path("resources").get(kind);
        if (rule == null) throw new IllegalArgumentException("unknown kind: " + kind);
        return rule;
    }
    public static void validateResource(String kind, String desired, boolean deleting, boolean agentWork) {
        JsonNode rule = resource(kind);
        if (agentWork && !"agent".equals(rule.path("owner").asText()))
            throw new IllegalArgumentException("server-managed kind cannot be agent work");
        if (!contains(rule.path("desiredStates"), desired)) throw new IllegalArgumentException("invalid desired state");
        boolean terminal = contains(rule.path("terminalStates"), rule.path("acknowledges").path(desired).asText());
        if (deleting != terminal) throw new IllegalArgumentException("deleting contradicts desired state");
    }
    public static boolean canTransition(String kind, String from, String to) {
        JsonNode rule;
        try { rule = resource(kind); } catch (IllegalArgumentException e) { return false; }
        if (!contains(rule.path("observedStates"), from) || !contains(rule.path("observedStates"), to)) return false;
        if (from.equals(to)) return true;
        var queue = new ArrayDeque<String>(); var seen = new HashSet<String>();
        queue.add(from); seen.add(from);
        while (!queue.isEmpty()) {
            String current = queue.remove();
            for (JsonNode next : rule.path("transitions").path(current)) {
                String state = next.asText();
                if (state.equals(to)) return true;
                if (seen.add(state)) queue.add(state);
            }
        }
        return false;
    }
    /** false = stale generation ignored; null generation never acknowledges work. */
    public static boolean validateObservation(String kind, String desired, String previous, String state,
            long targetGeneration, Long observedGeneration) {
        return validateObservationForOwner(kind, desired, previous, state, targetGeneration, observedGeneration, "agent");
    }
    public static boolean validateServerObservation(String kind, String desired, String previous, String state,
            long targetGeneration, Long observedGeneration) {
        return validateObservationForOwner(kind, desired, previous, state, targetGeneration, observedGeneration, "server");
    }
    private static boolean validateObservationForOwner(String kind, String desired, String previous, String state,
            long targetGeneration, Long observedGeneration, String owner) {
        JsonNode rule = resource(kind);
        if (!owner.equals(rule.path("owner").asText())) throw new IllegalArgumentException("invalid observation owner/kind");
        if (!contains(rule.path("desiredStates"), desired) || !contains(rule.path("observedStates"), state))
            throw new IllegalArgumentException("invalid lifecycle state");
        if (targetGeneration < 1) throw new IllegalArgumentException("invalid target generation");
        if (observedGeneration != null) {
            if (observedGeneration < 1 || observedGeneration > targetGeneration)
                throw new IllegalArgumentException("invalid or future generation");
            if (observedGeneration < targetGeneration) return false;
            if (!state.equals(rule.path("acknowledges").path(desired).asText()))
                throw new IllegalArgumentException("generation acknowledged before convergence");
        }
        if (("deleting".equals(state) || contains(rule.path("terminalStates"), state))
                && !contains(rule.path("terminalStates"), rule.path("acknowledges").path(desired).asText()))
            throw new IllegalArgumentException("deletion observation without deletion intent");
        if (previous == null || previous.isEmpty()) previous = rule.path("initialState").asText();
        if (!canTransition(kind, previous, state)) throw new IllegalArgumentException("invalid lifecycle transition");
        return true;
    }
    public static boolean canEntityTransition(String entity, String from, String to) {
        if (entity == null) return false;
        JsonNode rule = CATALOG.path("entities").get(entity);
        if (rule == null || !contains(rule.path("states"), from) || !contains(rule.path("states"), to)) return false;
        return from.equals(to) || contains(rule.path("transitions").path(from), to);
    }
}

package com.ciagotech.kan.server.auth;

import org.junit.jupiter.api.Test;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.data.redis.core.ValueOperations;
import java.lang.reflect.Proxy;
import java.time.Duration;
import java.util.concurrent.atomic.AtomicReference;
import static org.junit.jupiter.api.Assertions.*;

class ValkeyNonceStoreTest {
    @SuppressWarnings("unchecked")
    private StringRedisTemplate template(Boolean answer, AtomicReference<Object[]> captured) {
        ValueOperations<String,String> ops = (ValueOperations<String,String>) Proxy.newProxyInstance(
                ValueOperations.class.getClassLoader(), new Class[]{ValueOperations.class},(proxy,method,args) -> {
                    assertEquals("setIfAbsent",method.getName());captured.set(args);return answer;
                });
        return new StringRedisTemplate() {
            @Override public ValueOperations<String,String> opsForValue() { return ops; }
        };
    }
    @Test void reservesAtomicallyWithTtlAndSafeKeyEncoding() {
        var captured = new AtomicReference<Object[]>();
        assertTrue(new ValkeyNonceStore(template(true,captured)).checkAndStore("key:1","nonce:2",600));
        assertTrue(((String)captured.get()[0]).startsWith("kan:nonce:"));
        assertEquals("1",captured.get()[1]);assertEquals(Duration.ofSeconds(600),captured.get()[2]);
        String first = (String)captured.get()[0];
        new ValkeyNonceStore(template(true,captured)).checkAndStore("key","1:nonce:2",600);
        assertNotEquals(first,captured.get()[0]);
        assertFalse(new ValkeyNonceStore(template(false,captured)).checkAndStore("key","nonce",600));
    }
    @Test void missingReplyAndInvalidTtlFailClosed() {
        var store = new ValkeyNonceStore(template(null,new AtomicReference<>()));
        assertThrows(IllegalStateException.class,() -> store.checkAndStore("key","nonce",600));
        assertThrows(IllegalArgumentException.class,() -> store.checkAndStore("key","nonce",0));
    }
}

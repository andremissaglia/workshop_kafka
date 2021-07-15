package org.example.aux;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.kafka.common.errors.SerializationException;
import org.apache.kafka.common.serialization.Deserializer;
import org.apache.kafka.common.serialization.Serde;
import org.apache.kafka.common.serialization.Serializer;
import org.example.CountAndSum;

import java.io.IOException;
import java.util.Map;

public class JSONSerde implements Serializer<CountAndSum>, Deserializer<CountAndSum>, Serde<CountAndSum> {
    private static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();

    @Override
    public CountAndSum deserialize(String topic, byte[] data) {
        if (data == null) {
            return null;
        }

        try {
            return OBJECT_MAPPER.readValue(data, CountAndSum.class);
        } catch (final IOException e) {
            throw new SerializationException(e);
        }
    }

    @Override
    public byte[] serialize(String topic, CountAndSum data) {
        if (data == null) {
            return null;
        }

        try {
            return OBJECT_MAPPER.writeValueAsBytes(data);
        } catch (final Exception e) {
            throw new SerializationException("Error serializing JSON message", e);
        }
    }

    @Override
    public void close() {}

    @Override
    public Serializer<CountAndSum> serializer() {
        return this;
    }

    @Override
    public Deserializer<CountAndSum> deserializer() {
        return this;
    }

    @Override
    public void configure(Map<String, ?> configs, boolean isKey) {}
}

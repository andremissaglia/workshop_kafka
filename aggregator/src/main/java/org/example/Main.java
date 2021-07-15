package org.example;

import org.apache.kafka.common.serialization.Serde;
import org.apache.kafka.common.serialization.Serdes;
import org.apache.kafka.streams.KafkaStreams;
import org.apache.kafka.streams.StreamsBuilder;
import org.apache.kafka.streams.StreamsConfig;
import org.apache.kafka.streams.Topology;
import org.apache.kafka.streams.kstream.*;
import org.example.aux.JSONSerde;

import java.util.Properties;


public class Main {

    public static void main(String[] args) {
        // Serializers + Deserializers
        Serde<String> stringSerde = Serdes.String();
        Serde<CountAndSum> countAndSumSerde = new JSONSerde();

        // Configs
        String kafkaBrokers = System.getenv().getOrDefault("KAFKA_BROKERS", "localhost:9092");
        String sourceTopic = "votes";
        String destinationTopic = "ratings";

        Properties settings = new Properties();
        settings.put(StreamsConfig.APPLICATION_ID_CONFIG, "vote-aggregator");
        settings.put(StreamsConfig.BOOTSTRAP_SERVERS_CONFIG, kafkaBrokers);
        settings.put(StreamsConfig.PROCESSING_GUARANTEE_CONFIG, "exactly_once");

        // Building the stream
        StreamsBuilder builder = new StreamsBuilder();

        KStream<String, String> textLines = builder.stream(
                sourceTopic,
                Consumed.with(stringSerde, stringSerde)
        );

        textLines.mapValues((ValueMapper<String, Integer>) Integer::parseInt)
                .groupByKey()
                .aggregate(
                        () -> new CountAndSum(0, 0),
                        (key, value, aggregate) -> new CountAndSum(aggregate.sum + value, aggregate.count + 1),
                        Materialized.with(stringSerde, countAndSumSerde)
                )
                .toStream()
                .mapValues((countAndSum) -> String.valueOf((float) countAndSum.sum / countAndSum.count))
                .to(
                        destinationTopic,
                        Produced.with(stringSerde, stringSerde)
                )
        ;

        // Finish building and start executing
        Topology topology =  builder.build();
        KafkaStreams streams = new KafkaStreams(topology, settings);
        streams.start();
    }
}

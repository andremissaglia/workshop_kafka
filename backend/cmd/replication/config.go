package main

var config Config

type Config struct {
	Kafka struct {
		Brokers []string `default:"localhost:9092"`
		Topic   string   `default:"ratings"`
		GroupID string   `default:"replication-worker"`
	}
	DB struct {
		Host string `default:"db"`
		User string `default:"postgres"`
		Pass string `default:"q1w2e3r4"`
		Name string `default:"postgres"`
	}
}

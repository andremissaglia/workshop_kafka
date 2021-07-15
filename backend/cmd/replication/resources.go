package main

import (
	"database/sql"
	"fmt"

	"github.com/andremissaglia/workshop_kafka/backend/service/replication"
	"github.com/andremissaglia/workshop_kafka/backend/service/replication/implreplication"

	_ "github.com/lib/pq"
)

var (
	worker replication.Worker

	storeRatingGateway replication.StoreRatingGateway
	postgresDB         *sql.DB
)

func getWorker() replication.Worker {
	if worker == nil {
		worker = replication.NewWorker(
			config.Kafka.Brokers,
			config.Kafka.Topic,
			config.Kafka.GroupID,
			getStoreRatingGateway(),
		)
	}
	return worker
}

func getStoreRatingGateway() replication.StoreRatingGateway {
	if storeRatingGateway == nil {
		storeRatingGateway = implreplication.NewStoreRatingPostgresAdapter(getPostgresDB())
	}
	return storeRatingGateway
}

func getPostgresDB() *sql.DB {
	if postgresDB == nil {
		connStr := fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			config.DB.User,
			config.DB.Pass,
			config.DB.Host,
			config.DB.Name,
		)

		var err error
		postgresDB, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
	}
	return postgresDB
}

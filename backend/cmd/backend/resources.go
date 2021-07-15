package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/andremissaglia/workshop_kafka/backend/service/cat"
	"github.com/andremissaglia/workshop_kafka/backend/service/cat/apicat"
	"github.com/andremissaglia/workshop_kafka/backend/service/cat/implcat"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	httpServer *http.Server

	voteEndpoint     endpoint.Endpoint
	listCatsEndpoint endpoint.Endpoint

	catService cat.Service

	listCatsGateway cat.ListCatsGateway
	voteGateway     cat.VoteGateway
	postgresDB      *sql.DB
)

func getHTTPServer() *http.Server {
	if httpServer == nil {
		r := mux.NewRouter()

		getCatsHandler := apicat.MakeListCatsHTTPHandlers(getListCatsEndpoint())
		voteHandler := apicat.MakeVoteHTTPHandlers(getVoteEndpoint())
		r.Handle("/listCats", getCatsHandler).Methods(http.MethodGet, http.MethodOptions)
		r.Handle("/vote", voteHandler).Methods(http.MethodPost, http.MethodOptions)
		r.Use(apicat.CORSMiddleware)
		httpServer = &http.Server{
			Addr:    ":8000",
			Handler: r,
		}
	}
	return httpServer

}

func getVoteEndpoint() endpoint.Endpoint {
	if voteEndpoint == nil {
		voteEndpoint = apicat.MakeVoteEndpoint(getCatService())
	}
	return voteEndpoint
}

func getListCatsEndpoint() endpoint.Endpoint {
	if listCatsEndpoint == nil {
		listCatsEndpoint = apicat.MakeListCatsEndpoint(getCatService())
	}
	return listCatsEndpoint
}

func getCatService() cat.Service {
	if catService == nil {
		catService = cat.NewService(
			getListCatsGateway(),
			getVoteGateway(),
		)
	}
	return catService
}

func getListCatsGateway() cat.ListCatsGateway {
	if listCatsGateway == nil {
		listCatsGateway = implcat.NewListCatsPostgresAdapter(getPostgresDB())
	}
	return listCatsGateway
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

func getVoteGateway() cat.VoteGateway {
	if voteGateway == nil {
		voteGateway = implcat.NewVoteKafkaAdapter(config.Kafka.Brokers, config.Kafka.Topic)
	}
	return voteGateway
}

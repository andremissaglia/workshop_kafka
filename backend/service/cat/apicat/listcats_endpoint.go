package apicat

import (
	"context"

	"github.com/andremissaglia/workshop_kafka/backend/service/cat"
	"github.com/go-kit/kit/endpoint"
)

type ListCatsEndpointResponse struct {
	Cats []cat.Cat
}

func MakeListCatsEndpoint(catService cat.Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		cats, err := catService.ListCats(ctx)
		if err != nil {
			return nil, err
		}
		return ListCatsEndpointResponse{
			Cats: cats,
		}, nil
	}
}

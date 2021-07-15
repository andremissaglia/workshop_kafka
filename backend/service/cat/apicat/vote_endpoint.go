package apicat

import (
	"context"

	"github.com/andremissaglia/workshop_kafka/backend/service/cat"
	"github.com/go-kit/kit/endpoint"
)

type VoteEndpointRequest struct {
	CatID int
	Vote  int
}

type VoteEndpointResponse struct {
	Success bool
}

func MakeVoteEndpoint(catService cat.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VoteEndpointRequest)
		err := catService.Vote(ctx, req.CatID, req.Vote)
		if err != nil {
			return nil, err
		}
		return VoteEndpointResponse{
			Success: true,
		}, nil
	}
}

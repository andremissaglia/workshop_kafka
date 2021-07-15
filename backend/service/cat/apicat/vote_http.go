package apicat

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeVoteHTTPHandlers(e endpoint.Endpoint) http.Handler {
	return kithttp.NewServer(
		e,
		decodeVoteRequest,
		encodeVoteResponse,
	)
}
func decodeVoteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var endpointRequest VoteEndpointRequest

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&endpointRequest)
	if err != nil {
		return nil, err
	}

	return endpointRequest, nil
}
func encodeVoteResponse(ctx context.Context, w http.ResponseWriter, r interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	endpointResponse := r.(VoteEndpointResponse)
	return json.NewEncoder(w).Encode(endpointResponse)
}

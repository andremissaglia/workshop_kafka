package apicat

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeListCatsHTTPHandlers(e endpoint.Endpoint) http.Handler {
	return kithttp.NewServer(
		e,
		decodeListCatsRequest,
		encodeListCatsResponse,
	)
}
func decodeListCatsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
func encodeListCatsResponse(ctx context.Context, w http.ResponseWriter, r interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	endpointResponse := r.(ListCatsEndpointResponse)
	return json.NewEncoder(w).Encode(endpointResponse)
}

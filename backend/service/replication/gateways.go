package replication

import "context"

type StoreRatingGateway interface {
	Store(ctx context.Context, catID int, rating float32) error
}

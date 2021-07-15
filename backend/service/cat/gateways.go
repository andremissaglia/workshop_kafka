package cat

import "context"

type ListCatsGateway interface {
	ListCats(ctx context.Context) ([]Cat, error)
}

type VoteGateway interface {
	Vote(ctx context.Context, catID int, vote int) error
}

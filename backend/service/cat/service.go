package cat

import (
	"context"
	"fmt"
)

type Service interface {
	Vote(ctx context.Context, catID int, vote int) error
	ListCats(ctx context.Context) ([]Cat, error)
}

type service struct {
	listCatsGateway ListCatsGateway
	voteGateway     VoteGateway
}

func NewService(
	listCatsGateway ListCatsGateway,
	voteGateway VoteGateway,
) Service {
	return &service{
		listCatsGateway: listCatsGateway,
		voteGateway:     voteGateway,
	}
}

func (s *service) Vote(ctx context.Context, catID int, vote int) error {
	fmt.Printf("Votando para %d com nota %d\n", catID, vote)
	err := s.voteGateway.Vote(ctx, catID, vote)
	if err != nil {
		fmt.Println("Erro ao votar: ", err)
	}
	return err
}

func (s *service) ListCats(ctx context.Context) ([]Cat, error) {
	return s.listCatsGateway.ListCats(ctx)
}

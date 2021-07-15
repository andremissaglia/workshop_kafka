package implreplication

import (
	"context"
	"database/sql"

	"github.com/andremissaglia/workshop_kafka/backend/service/replication"
)

type storeRatingPostgresAdapter struct {
	db *sql.DB
}

func NewStoreRatingPostgresAdapter(db *sql.DB) replication.StoreRatingGateway {
	return &storeRatingPostgresAdapter{
		db: db,
	}
}

func (a *storeRatingPostgresAdapter) Store(ctx context.Context, catID int, rating float32) error {
	stmt, err := a.db.Prepare("UPDATE cats SET rating = $2 WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, catID, rating)
	return err

}

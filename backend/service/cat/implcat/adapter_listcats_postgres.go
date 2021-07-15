package implcat

import (
	"context"
	"database/sql"

	"github.com/andremissaglia/workshop_kafka/backend/service/cat"
)

type listCatsPostgresAdapter struct {
	db *sql.DB
}

func NewListCatsPostgresAdapter(db *sql.DB) cat.ListCatsGateway {
	return &listCatsPostgresAdapter{
		db: db,
	}
}

func (a *listCatsPostgresAdapter) ListCats(ctx context.Context) ([]cat.Cat, error) {
	rows, err := a.db.Query("SELECT ID, Name, PhotoURL, Rating FROM cats ORDER BY ID;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []cat.Cat{}
	var current cat.Cat

	for rows.Next() {
		err = rows.Scan(&current.ID, &current.Name, &current.PhotoURL, &current.Rating)
		if err != nil {
			return nil, err
		}
		result = append(result, current)
	}
	return result, nil
}

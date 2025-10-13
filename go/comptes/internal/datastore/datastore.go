package datastore

import (
	"comptes/internal/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Datastore struct {
	dbpool *pgxpool.Pool
}

func NewDatastore(cfg *config.Config) (*Datastore, error) {
	dbpool, err := pgxpool.New(context.Background(), cfg.Db.Connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	// defer dbpool.Close()

	return &Datastore{
		dbpool: dbpool,
	}, nil
}

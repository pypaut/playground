package internal

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Datastore struct {
	dbpool *pgxpool.Pool
}

func NewDatastore(cfg *Config) (*Datastore, error) {
	dbpool, err := pgxpool.New(context.Background(), cfg.Db.Connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	// defer dbpool.Close()

	return &Datastore{
		dbpool: dbpool,
	}, nil
}

func (d *Datastore) ListBudgets() (budgets []*Budget, err error) {
	rows, err := d.dbpool.Query(context.Background(), "select * from budgets")
	if err != nil {
		return nil, fmt.Errorf("could not list budgets: %w", err)
	}

	for rows.Next() {
		var budget Budget
		err := rows.Scan(
			&budget.Tag,
			&budget.Amount,
			&budget.Date,
			&budget.Label,
		)

		if err != nil {
			return nil, fmt.Errorf("could not scan budget: %w", err)
		}

		budgets = append(budgets, &budget)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate budgets: %w", err)
	}

	return
}

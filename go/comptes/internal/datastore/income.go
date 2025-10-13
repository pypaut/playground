package datastore

import (
	"comptes/internal/model"
	"context"
	"fmt"
)

func (d *Datastore) ListIncomes(year, month int) (incomes []*model.Income, err error) {
	rows, err := d.dbpool.Query(
		context.Background(),
		"select * from incomes where extract(year from date) = $1 and extract(month from date) = $2", year, month,
	)
	if err != nil {
		return nil, fmt.Errorf("could not list incomes: %w", err)
	}

	for rows.Next() {
		var income model.Income
		err := rows.Scan(
			&income.ID,
			&income.Label,
			&income.Amount,
			&income.Date,
		)

		if err != nil {
			return nil, fmt.Errorf("could not scan income: %w", err)
		}

		incomes = append(incomes, &income)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate incomes: %w", err)
	}

	return
}

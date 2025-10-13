package datastore

import (
	"comptes/internal/model"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

func (d *Datastore) AddBudget(budget *model.Budget) error {
	query := `INSERT INTO budgets (label, amount, date, tag_id)
VALUES (@label, @amount, @date, @tagId)`
	args := pgx.NamedArgs{
		"label":  budget.Label,
		"amount": budget.Amount,
		"date":   budget.Date,
		"tagId":  budget.TagID.String(),
	}

	_, err := d.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("error creating budget: %s", err)
	}

	return nil
}

func (d *Datastore) ListBudgets(year, month int) (budgets []*model.Budget, err error) {
	rows, err := d.dbpool.Query(
		context.Background(),
		"select * from budgets where extract(year from date) = $1 and extract(month from date) = $2",
		year,
		month,
	)
	if err != nil {
		return nil, fmt.Errorf("could not list budgets: %w", err)
	}

	for rows.Next() {
		var budget model.Budget
		err := rows.Scan(
			&budget.ID,
			&budget.Label,
			&budget.Amount,
			&budget.Date,
			&budget.TagID,
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

func (d *Datastore) ListBudgetsForTagId(tagId uuid.UUID, year, month int) (budgets []*model.Budget, err error) {
	rows, err := d.dbpool.Query(
		context.Background(),
		"select * from budgets where tag_id=$1 and extract(year from date) = $2 and extract(month from date) = $3",
		tagId.String(),
		year,
		month,
	)
	if err != nil {
		return nil, fmt.Errorf("could not list budgets: %w", err)
	}

	for rows.Next() {
		var budget model.Budget
		err := rows.Scan(
			&budget.ID,
			&budget.Label,
			&budget.Amount,
			&budget.Date,
			&budget.TagID,
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

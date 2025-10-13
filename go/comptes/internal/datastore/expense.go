package datastore

import (
	"comptes/internal/model"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
)

func (d *Datastore) ListExpenses(year, month int) (expenses []*model.Expense, err error) {
	rows, err := d.dbpool.Query(
		context.Background(),
		"select * from expenses where extract(year from date) = $1 and extract(month from date) = $2",
		year,
		month,
	)
	if err != nil {
		return nil, fmt.Errorf("could not list expenses: %w", err)
	}

	for rows.Next() {
		var expense model.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.Label,
			&expense.Amount,
			&expense.Date,
			&expense.BudgetID,
		)

		if err != nil {
			return nil, fmt.Errorf("could not scan expense: %w", err)
		}

		expenses = append(expenses, &expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate expenses: %w", err)
	}

	return
}

func (d *Datastore) ListExpensesForBudget(year, month int, budgetId uuid.UUID) (expenses []*model.Expense, err error) {
	rows, err := d.dbpool.Query(
		context.Background(),
		"select * from expenses where budget_id=$1 and extract(year from date) = $2 and extract(month from date) = $3",
		budgetId,
		year,
		month,
	)
	if err != nil {
		return nil, fmt.Errorf("could not list expenses: %w", err)
	}

	for rows.Next() {
		var expense model.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.Label,
			&expense.Amount,
			&expense.Date,
			&expense.BudgetID,
		)

		if err != nil {
			return nil, fmt.Errorf("could not scan expense: %w", err)
		}

		expenses = append(expenses, &expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate expenses: %w", err)
	}

	return
}

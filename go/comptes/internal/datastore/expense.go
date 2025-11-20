package datastore

import (
	"comptes/internal/model"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
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

func (d *Datastore) AddExpense(expense *model.Expense) error {
	query := `INSERT INTO expenses (label, amount, date, budget_id)
VALUES (@label, @amount, @date, @budgetId)`
	args := pgx.NamedArgs{
		"label":    expense.Label,
		"amount":   expense.Amount,
		"date":     expense.Date,
		"budgetId": expense.BudgetID.String(),
	}

	_, err := d.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("error creating expense: %s", err)
	}

	return nil
}

func (d *Datastore) RemoveExpense(expenseId uuid.UUID) error {
	query := `DELETE FROM expenses WHERE id=@expenseId`
	args := pgx.NamedArgs{
		"expenseId": expenseId,
	}

	_, err := d.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("error removing expense: %s", err)
	}

	return nil
}

func (d *Datastore) GetExpense(expenseId uuid.UUID) (*model.Expense, error) {
	var expense model.Expense
	err := d.dbpool.QueryRow(context.Background(), "select * from expenses where id = $1", expenseId).Scan(
		&expense.ID,
		&expense.Label,
		&expense.Amount,
		&expense.Date,
		&expense.BudgetID,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get expense: %w", err)
	}

	return &expense, nil
}

func (d *Datastore) UpdateExpense(expense *model.Expense) error {
	query := `UPDATE expenses SET label=@label, amount=@amount, date=@date, budget_id=@budgetId WHERE id=@id`
	args := pgx.NamedArgs{
		"id":       expense.ID.String(),
		"label":    expense.Label,
		"amount":   expense.Amount,
		"date":     expense.Date,
		"budgetId": expense.BudgetID.String(),
	}

	_, err := d.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("error updating expense: %s", err)
	}

	return nil
}

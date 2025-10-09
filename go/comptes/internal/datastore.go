package internal

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
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

func (d *Datastore) AddBudget(budget *Budget) error {
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
		// return err
	}

	return nil
}

func (d *Datastore) ListBudgets(year, month int) (budgets []*Budget, err error) {
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
		var budget Budget
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

func (d *Datastore) GetTagByLabel(tagLabel string) (*Tag, error) {
	var tag Tag
	err := d.dbpool.QueryRow(context.Background(), "select * from tags where label = $1", tagLabel).Scan(&tag)
	if err != nil {
		return nil, fmt.Errorf("could not get tag: %w", err)
	}

	return &tag, nil
}

func (d *Datastore) ListBudgetsForTagId(tagId uuid.UUID, year, month int) (budgets []*Budget, err error) {
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
		var budget Budget
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

func (d *Datastore) ListTags() (tags []*Tag, err error) {
	rows, err := d.dbpool.Query(context.Background(), "select * from tags")
	if err != nil {
		return nil, fmt.Errorf("could not list tags: %w", err)
	}

	for rows.Next() {
		var tag Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Label,
			&tag.Description,
			&tag.Icon,
		)

		if err != nil {
			return nil, fmt.Errorf("could not scan tags: %w", err)
		}

		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate tags: %w", err)
	}

	return
}

func (d *Datastore) ListExpenses(year, month int) (expenses []*Expense, err error) {
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
		var expense Expense
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

func (d *Datastore) ListExpensesForBudget(year, month int, budget string) (expenses []*Expense, err error) {
	rows, err := d.dbpool.Query(
		context.Background(),
		"select * from expenses where budget=$1 and extract(year from date) = $2 and extract(month from date) = $3",
		budget,
		year,
		month,
	)
	if err != nil {
		return nil, fmt.Errorf("could not list expenses: %w", err)
	}

	for rows.Next() {
		var expense Expense
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

func (d *Datastore) ListIncomes(year, month int) (incomes []*Income, err error) {
	rows, err := d.dbpool.Query(
		context.Background(),
		"select * from incomes where extract(year from date) = $1 and extract(month from date) = $2", year, month,
	)
	if err != nil {
		return nil, fmt.Errorf("could not list incomes: %w", err)
	}

	for rows.Next() {
		var income Income
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

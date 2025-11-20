package datastore

import (
	"comptes/internal/model"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
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

func (d *Datastore) GetIncome(incomeId uuid.UUID) (*model.Income, error) {
	var income model.Income
	err := d.dbpool.QueryRow(context.Background(), "select * from incomes where id = $1", incomeId).Scan(
		&income.ID,
		&income.Label,
		&income.Amount,
		&income.Date,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get income: %w", err)
	}

	return &income, nil
}

func (d *Datastore) AddIncome(income *model.Income) error {
	query := `INSERT INTO incomes (label, amount, date)
VALUES (@label, @amount, @date)`
	args := pgx.NamedArgs{
		"label":  income.Label,
		"amount": income.Amount,
		"date":   income.Date,
	}

	_, err := d.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("error creating income: %s", err)
	}

	return nil
}

func (d *Datastore) RemoveIncome(incomeId uuid.UUID) error {
	query := `DELETE FROM incomes WHERE id=@incomeId`
	args := pgx.NamedArgs{
		"incomeId": incomeId,
	}

	_, err := d.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("error removing income: %s", err)
	}

	return nil
}

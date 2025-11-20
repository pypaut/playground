package cli

import (
	"comptes/internal/datastore"
	"comptes/internal/table"
)

type ExpensesCmd struct {
	datastore *datastore.Datastore

	Year  int
	Month int
}

func (cmd *ExpensesCmd) Run() error {
	year, month := parseDate(cmd.Year, cmd.Month)
	expensesTable := table.BuildExpensesTable(cmd.datastore, year, month)
	expensesTable.Render()

	return nil
}

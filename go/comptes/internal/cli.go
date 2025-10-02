package internal

import (
	"fmt"
	"time"
)

type CLI struct {
	Budgets  BudgetsCmd  `cmd:"budgets" help:"Show budgets"`
	Expenses ExpensesCmd `cmd:"accounts" help:"Show expenses"`
}

func NewCli(datastore *Datastore) *CLI {
	return &CLI{
		Budgets: BudgetsCmd{
			datastore: datastore,
		},
		Expenses: ExpensesCmd{
			datastore: datastore,
		},
	}
}

type BudgetsCmd struct {
	datastore *Datastore

	Year  int
	Month int
}

func (cmd *BudgetsCmd) Run() error {
	date := parseDate(cmd.Year, cmd.Month)
	fmt.Printf("Date: %v\n", date)

	incomesTable := BuildIncomesTable(cmd.datastore)
	budgetsTables := BuildBudgetsTables(cmd.datastore)
	remainTable := BuildRemainTable(cmd.datastore)

	incomesTable.Render()
	for _, t := range budgetsTables {
		t.Render()
	}
	remainTable.Render()

	return nil
}

type ExpensesCmd struct {
	datastore *Datastore
}

func (cmd *ExpensesCmd) Run() error {
	expensesTable := BuildExpensesTable(cmd.datastore)
	expensesTable.Render()
	return nil
}

func parseDate(year int, month int) string {
	now := time.Now()
	if year == 0 {
		year = now.Year()
	}

	if month == 0 {
		month = int(now.Month())
	}

	return time.Date(
		year, time.Month(month), 1, 0, 0, 0, 0, time.Local,
	).Format("2006-01")
}

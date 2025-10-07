package internal

import (
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

	Year        int
	Month       int
	Proportions bool
}

func (cmd *BudgetsCmd) Run() error {
	year, month := parseDate(cmd.Year, cmd.Month)

	incomesTable := BuildIncomesTable(cmd.datastore, year, month)
	budgetsTables := BuildBudgetsTables(cmd.datastore, year, month)
	remainTable := BuildRemainTable(cmd.datastore, year, month)

	incomesTable.Render()
	for _, t := range budgetsTables {
		t.Render()
	}
	if cmd.Proportions {
		proportionsTable := BuildProportionsTable(cmd.datastore, year, month)
		proportionsTable.Render()
	}
	remainTable.Render()

	return nil
}

type ExpensesCmd struct {
	datastore *Datastore

	Year  int
	Month int
}

func (cmd *ExpensesCmd) Run() error {
	year, month := parseDate(cmd.Year, cmd.Month)
	expensesTable := BuildExpensesTable(cmd.datastore, year, month)
	expensesTable.Render()
	return nil
}

func parseDate(year int, month int) (int, int) {
	now := time.Now()
	if year == 0 || month == 0 {
		year = now.Year()
		month = int(now.Month())
	}

	return year, month

	// return time.Date(
	// 	year, time.Month(month), 1, 0, 0, 0, 0, time.Local,
	// ) // .Format("2006-01")
}

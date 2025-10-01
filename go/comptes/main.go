package main

import (
	"comptes/internal"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	cfg, err := internal.LoadConfig("config.yml")
	if err != nil {
		panic(err)
	}

	datastore, err := internal.NewDatastore(cfg)
	if err != nil {
		panic(err)
	}

	incomesTable := buildIncomesTable(datastore)
	budgetsTable := buildBudgetsTable(datastore)
	expensesTable := buildExpensesTable(datastore)

	incomesTable.Render()
	budgetsTable.Render()
	expensesTable.Render()
}

func buildIncomesTable(datastore *internal.Datastore) (incomesTable table.Writer) {
	incomesTable = table.NewWriter()
	incomesTable.SetOutputMirror(os.Stdout)

	incomes, err := datastore.ListIncomes()
	if err != nil {
		panic(err)
	}

	incomesTable.AppendHeader(table.Row{"Income", "Amount"})

	for _, income := range incomes {
		incomesTable.AppendRow(table.Row{income.Label, income.Amount})
	}

	return
}

func buildBudgetsTable(datastore *internal.Datastore) (budgetsTable table.Writer) {
	budgetsTable = table.NewWriter()
	budgetsTable.SetOutputMirror(os.Stdout)

	budgets, err := datastore.ListBudgets()
	if err != nil {
		panic(err)
	}

	budgetsTable.AppendHeader(table.Row{"Budget", "Amount"})

	for _, budget := range budgets {
		budgetsTable.AppendRow(table.Row{budget.Label, budget.Amount})
	}

	return
}

func buildExpensesTable(datastore *internal.Datastore) (expensesTable table.Writer) {
	expensesTable = table.NewWriter()
	expensesTable.SetOutputMirror(os.Stdout)

	expenses, err := datastore.ListExpenses()
	if err != nil {
		panic(err)
	}

	expensesTable.AppendHeader(table.Row{"Expense", "Amount", "Date"})

	for _, expense := range expenses {
		expensesTable.AppendRow(table.Row{expense.Label, expense.Amount, expense.Date})
	}

	return
}

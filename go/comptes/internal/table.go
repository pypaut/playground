package internal

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func BuildBudgetsTables(datastore *Datastore) (budgetsTables []table.Writer) {
	tags, err := datastore.ListTags()
	if err != nil {
		panic(err)
	}

	for _, tag := range tags {
		budgetsTables = append(budgetsTables, BuildBudgetsTableForTag(datastore, tag.Label))
	}

	return
}

func BuildIncomesTable(datastore *Datastore) (incomesTable table.Writer) {
	incomesTable = table.NewWriter()
	incomesTable.SetOutputMirror(os.Stdout)

	incomes, err := datastore.ListIncomes()
	if err != nil {
		panic(err)
	}

	incomesTable.AppendHeader(table.Row{"Income", "Amount"})

	for _, income := range incomes {
		incomesTable.AppendRow(table.Row{income.Label, income.Amount / 100})
	}

	return
}

func BuildBudgetsTable(datastore *Datastore) (budgetsTable table.Writer) {
	budgetsTable = table.NewWriter()
	budgetsTable.SetOutputMirror(os.Stdout)

	budgets, err := datastore.ListBudgets()
	if err != nil {
		panic(err)
	}

	budgetsTable.AppendHeader(table.Row{"Budget", "Amount"})

	for _, budget := range budgets {
		budgetsTable.AppendRow(table.Row{budget.Label, budget.Amount / 100})
	}

	return
}

func BuildBudgetsTableForTag(
	datastore *Datastore, tagLabel string,
) (budgetsTable table.Writer) {
	budgetsTable = table.NewWriter()
	budgetsTable.SetOutputMirror(os.Stdout)

	budgets, err := datastore.ListBudgetsForTag(tagLabel)
	if err != nil {
		panic(err)
	}

	budgetsTable.AppendHeader(table.Row{tagLabel, "Budget", "Amount"})

	for _, budget := range budgets {
		budgetsTable.AppendRow(table.Row{"", budget.Label, budget.Amount / 100})
	}

	return
}

func BuildExpensesTable(datastore *Datastore) (expensesTable table.Writer) {
	expensesTable = table.NewWriter()
	expensesTable.SetOutputMirror(os.Stdout)

	expenses, err := datastore.ListExpenses()
	if err != nil {
		panic(err)
	}

	expensesTable.AppendHeader(table.Row{"Expense", "Amount", "Date"})

	for _, expense := range expenses {
		expensesTable.AppendRow(table.Row{expense.Label, expense.Amount / 100, expense.Date})
	}

	return
}

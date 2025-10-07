package internal

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func BuildBudgetsTables(datastore *Datastore, year, month int) (budgetsTables []table.Writer) {
	tags, err := datastore.ListTags()
	if err != nil {
		panic(err)
	}

	for _, tag := range tags {
		budgetsTables = append(budgetsTables, BuildBudgetsTableForTag(datastore, tag.Label, year, month))
	}

	return
}

func BuildIncomesTable(datastore *Datastore, year, month int) (incomesTable table.Writer) {
	incomesTable = table.NewWriter()
	incomesTable.SetOutputMirror(os.Stdout)

	incomes, err := datastore.ListIncomes(year, month)
	if err != nil {
		panic(err)
	}

	incomesTable.AppendHeader(table.Row{"Income", "Amount"})

	for _, income := range incomes {
		incomesTable.AppendRow(table.Row{income.Label, income.Amount / 100})
	}

	return
}

func BuildBudgetsTable(datastore *Datastore, year, month int) (budgetsTable table.Writer) {
	budgetsTable = table.NewWriter()
	budgetsTable.SetOutputMirror(os.Stdout)

	budgets, err := datastore.ListBudgets(year, month)
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
	datastore *Datastore, tagLabel string, year, month int,
) (budgetsTable table.Writer) {
	budgetsTable = table.NewWriter()
	budgetsTable.SetOutputMirror(os.Stdout)

	budgets, err := datastore.ListBudgetsForTag(tagLabel, year, month)
	if err != nil {
		panic(err)
	}

	budgetsTable.AppendHeader(table.Row{tagLabel, "Budget", "Amount"})

	for _, budget := range budgets {
		budgetsTable.AppendRow(table.Row{"", budget.Label, budget.Amount / 100})
	}

	return
}

func BuildExpensesTable(datastore *Datastore, year, month int) (expensesTable table.Writer) {
	expensesTable = table.NewWriter()
	expensesTable.SetOutputMirror(os.Stdout)

	expenses, err := datastore.ListExpenses(year, month)
	if err != nil {
		panic(err)
	}

	expensesTable.AppendHeader(table.Row{"Date", "Expense", "Amount"})

	for _, expense := range expenses {
		expensesTable.AppendRow(table.Row{expense.Date.Format("2006-01-02"), expense.Label, expense.Amount / 100})
	}

	return
}

func BuildRemainTable(datastore *Datastore, year, month int) (remainTable table.Writer) {
	remainTable = table.NewWriter()
	remainTable.SetOutputMirror(os.Stdout)

	incomes, err := datastore.ListIncomes(year, month)
	if err != nil {
		panic(err)
	}

	totalIncomes := 0.

	for _, i := range incomes {
		totalIncomes += i.Amount
	}

	budgets, err := datastore.ListBudgets(year, month)
	if err != nil {
		panic(err)
	}

	totalBudgets := 0.

	for _, b := range budgets {
		totalBudgets += b.Amount
	}

	remain := totalIncomes - totalBudgets

	remainTable.AppendHeader(table.Row{"Remain"})
	remainTable.AppendRow(table.Row{remain / 100})

	return
}

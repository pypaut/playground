package main

import (
	"comptes/internal"
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

	incomesTable := internal.BuildIncomesTable(datastore)
	budgetsTables := internal.BuildBudgetsTables(datastore)
	expensesTable := internal.BuildExpensesTable(datastore)
	remainTable := internal.BuildRemainTable(datastore)

	incomesTable.Render()
	for _, t := range budgetsTables {
		t.Render()
	}

	remainTable.Render()
	expensesTable.Render()
}

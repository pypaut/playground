package cli

import (
	"comptes/internal/datastore"
	"time"
)

type CLI struct {
	Budgets  BudgetsCmd  `cmd:"budgets" help:"Manage budgets"`
	Expenses ExpensesCmd `cmd:"accounts" help:"Manage expenses"`
	Goose    GooseCmd    `cmd:"goose" help:"Manage Goose commands"`
}

func NewCli(datastore *datastore.Datastore) *CLI {
	return &CLI{
		Budgets: BudgetsCmd{
			datastore: datastore,
			List: BudgetsListCmd{
				datastore: datastore,
			},
			Add: BudgetsAddCmd{
				datastore: datastore,
			},
			Update: BudgetsUpdateCmd{
				datastore: datastore,
			},
			Remove: BudgetsRemoveCmd{
				datastore: datastore,
			},
		},
		Expenses: ExpensesCmd{
			datastore: datastore,
		},
		Goose: GooseCmd{
			datastore: datastore,
		},
	}
}

func parseDate(year int, month int) (int, int) {
	now := time.Now()
	if year == 0 || month == 0 {
		year = now.Year()
		month = int(now.Month())
	}

	return year, month
}

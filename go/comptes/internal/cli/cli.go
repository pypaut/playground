package cli

import (
	"comptes/internal/datastore"
	"comptes/internal/table"
	"database/sql"
	"time"

	"github.com/pressly/goose/v3"
)

type CLI struct {
	Budgets  BudgetsCmd  `cmd:"budgets" help:"Show budgets"`
	Expenses ExpensesCmd `cmd:"accounts" help:"Show expenses"`
	Goose    GooseCmd    `cmd:"goose" help:"Manage Goose commands"`
}

func NewCli(datastore *datastore.Datastore, sqlDb *sql.DB) *CLI {
	return &CLI{
		Budgets: BudgetsCmd{
			datastore: datastore,
		},
		Expenses: ExpensesCmd{
			datastore: datastore,
		},
		Goose: GooseCmd{
			db: sqlDb,
		},
	}
}

type BudgetsCmd struct {
	datastore *datastore.Datastore

	Year        int
	Month       int
	Proportions bool
}

func (cmd *BudgetsCmd) Run() error {
	year, month := parseDate(cmd.Year, cmd.Month)

	incomesTable := table.BuildIncomesTable(cmd.datastore, year, month)
	budgetsTables := table.BuildBudgetsTables(cmd.datastore, year, month)
	remainTable := table.BuildRemainTable(cmd.datastore, year, month)

	incomesTable.Render()
	for _, t := range budgetsTables {
		t.Render()
	}
	if cmd.Proportions {
		proportionsTable := table.BuildProportionsTable(cmd.datastore, year, month)
		proportionsTable.Render()
	}
	remainTable.Render()

	return nil
}

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

func parseDate(year int, month int) (int, int) {
	now := time.Now()
	if year == 0 || month == 0 {
		year = now.Year()
		month = int(now.Month())
	}

	return year, month
}

type GooseAction string

const (
	GooseActionUp     GooseAction = "up"
	GooseActionDown   GooseAction = "down"
	GooseActionStatus GooseAction = "status"
)

type GooseCmd struct {
	db *sql.DB

	Action GooseAction
}

func (cmd *GooseCmd) Run() error {
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	switch cmd.Action {
	case GooseActionUp:
		if err := goose.Up(cmd.db, "./migrations"); err != nil {
			panic(err)
		}

	case GooseActionDown:
		if err := goose.Down(cmd.db, "./migrations"); err != nil {
			panic(err)
		}

	case GooseActionStatus:
		if err := goose.Status(cmd.db, "./migrations"); err != nil {
			panic(err)
		}
	}

	return nil
}

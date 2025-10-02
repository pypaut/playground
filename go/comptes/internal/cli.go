package internal

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
}

func (cmd *BudgetsCmd) Run() error {
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

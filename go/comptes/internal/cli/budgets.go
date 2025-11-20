package cli

import (
	"comptes/internal/datastore"
	"comptes/internal/model"
	"comptes/internal/table"
	"fmt"
	"time"
)

type BudgetsCmd struct {
	datastore *datastore.Datastore

	Add    BudgetsAddCmd    `cmd:"add"`
	List   BudgetsListCmd   `cmd:"list"`
	Update BudgetsUpdateCmd `cmd:"update"`
	Remove BudgetsRemoveCmd `cmd:"rm"`
}

type BudgetsRemoveCmd struct {
	datastore *datastore.Datastore

	Label string `required:""`

	Year  int
	Month int
}

func (cmd *BudgetsRemoveCmd) Run() error {
	year, month := parseDate(cmd.Year, cmd.Month)

	budget, err := cmd.datastore.GetBudgetByLabelAndDate(cmd.Label, year, month)
	if err != nil {
		return err
	}

	return cmd.datastore.RemoveBudget(budget.ID)
}

type BudgetsUpdateCmd struct {
	datastore *datastore.Datastore

	Label  string  `required:""`
	Amount float64 `required:""`
	Tag    string  `required:""`

	Year  int
	Month int
}

func (cmd *BudgetsUpdateCmd) Run() error {
	year, month := parseDate(cmd.Year, cmd.Month)
	date, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%d-01", year, month))
	if err != nil {
		return nil
	}

	tagModel, err := cmd.datastore.GetTagByLabel(cmd.Tag)
	if err != nil {
		return err
	}

	budget, err := cmd.datastore.GetBudgetByLabelAndDate(cmd.Label, year, month)
	if err != nil {
		return err
	}

	return cmd.datastore.UpdateBudget(&model.Budget{
		ID:     budget.ID,
		Label:  cmd.Label,
		Amount: cmd.Amount * 100,
		Date:   date,
		TagID:  tagModel.ID,
	})
}

type BudgetsAddCmd struct {
	datastore *datastore.Datastore

	Label  string  `required:""`
	Amount float64 `required:""`
	Tag    string  `required:""`

	Year  int
	Month int
}

func (cmd *BudgetsAddCmd) Run() error {
	year, month := parseDate(cmd.Year, cmd.Month)
	date, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%d-01", year, month))
	if err != nil {
		return nil
	}

	tagModel, err := cmd.datastore.GetTagByLabel(cmd.Tag)
	if err != nil {
		return err
	}

	return cmd.datastore.AddBudget(&model.Budget{
		Label:  cmd.Label,
		Amount: cmd.Amount * 100,
		Date:   date,
		TagID:  tagModel.ID,
	})
}

type BudgetsListCmd struct {
	datastore *datastore.Datastore

	Year        int
	Month       int
	Proportions bool
}

func (cmd *BudgetsListCmd) Run() error {
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

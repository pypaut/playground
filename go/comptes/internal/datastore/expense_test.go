package datastore

import (
	"comptes/internal/config"
	"comptes/internal/model"
	"reflect"
	"testing"
	"time"
)

func TestListExpenses(t *testing.T) {
	// cfg, err := config.LoadConfig("../../config.yml")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// ds, err := NewDatastore(cfg)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	loadFixtures()

	expectedExpenses := []*model.Expense{
		{
			ID:       expenseLoyerUUID,
			Label:    "Loyer",
			Amount:   120000,
			Date:     time.Date(2025, 07, 2, 0, 0, 0, 0, time.UTC),
			BudgetID: budgetLoyerUUID,
		},
		{
			ID:       expenseLeclercUUID,
			Label:    "Leclerc",
			Amount:   4781,
			Date:     time.Date(2025, 07, 8, 0, 0, 0, 0, time.UTC),
			BudgetID: budgetCoursesUUID,
		},
	}

	expenses, err := ds.ListExpenses(2025, 7)
	if err != nil {
		t.Fatalf("ListExpenses: %v", err)
	}

	if len(expenses) != 2 {
		t.Fatalf("ListExpenses: got %d expenses, want 2", len(expenses))
	}

	for i := range expenses {
		if !reflect.DeepEqual(expenses[i], expectedExpenses[i]) {
			t.Fatalf("ListExpenses: got %v, want %v", expenses[i], expectedExpenses[i])
		}
	}
}

func TestAddExpenses(t *testing.T) {
	cfg, err := config.LoadConfig("../../config.yml")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := NewDatastore(cfg)
	if err != nil {
		t.Fatal(err)
	}

	testExpense := &model.Expense{
		Label:    "Repas Geoffrey travail",
		Amount:   150,
		Date:     time.Date(2025, 10, 9, 0, 0, 0, 0, time.UTC),
		BudgetID: budgetCoursesUUID,
	}

	err = ds.AddExpense(testExpense)
	if err != nil {
		t.Fatalf("AddBudget: %s", err)
	}

	expenses, err := ds.ListExpenses(2025, 10)
	if err != nil {
		t.Fatalf("ListBudgets: %v", err)
	}

	if len(expenses) != 1 {
		t.Fatalf("ListBudgets: got %d expenses, want 1", len(expenses))
	}

	if testExpense.Amount != expenses[0].Amount ||
		testExpense.Date != expenses[0].Date ||
		testExpense.BudgetID != expenses[0].BudgetID {
		t.Fatalf("ListExpenses: got %v, want %v", expenses[0], testExpense)
	}
}

func TestRemoveExpenses(t *testing.T) {
	cfg, err := config.LoadConfig("../../config.yml")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := NewDatastore(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = ds.RemoveExpense(expenseLoyerUUID)
	if err != nil {
		t.Fatalf("RemoveExpense: %s", err)
	}

	expenses, err := ds.ListExpenses(2025, 7)
	if err != nil {
		t.Fatalf("ListExpenses: %v", err)
	}

	if len(expenses) != 1 {
		t.Fatalf("ListExpenses: got %d expenses, want 1", len(expenses))
	}
}

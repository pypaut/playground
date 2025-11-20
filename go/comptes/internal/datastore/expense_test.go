package datastore

import (
	"comptes/internal/model"
	"reflect"
	"testing"
	"time"
)

func TestListExpenses(t *testing.T) {
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
	loadFixtures()

	testExpense := &model.Expense{
		Label:    "Repas Geoffrey travail",
		Amount:   150,
		Date:     time.Date(2025, 10, 9, 0, 0, 0, 0, time.UTC),
		BudgetID: budgetCoursesUUID,
	}

	err := ds.AddExpense(testExpense)
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
	loadFixtures()

	err := ds.RemoveExpense(expenseLoyerUUID)
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

func TestGetExpense(t *testing.T) {
	loadFixtures()

	expectedExpense := &model.Expense{
		ID:       expenseLoyerUUID,
		Label:    "Loyer",
		Amount:   120000,
		Date:     time.Date(2025, 07, 2, 0, 0, 0, 0, time.UTC),
		BudgetID: budgetLoyerUUID,
	}

	gotExpense, err := ds.GetExpense(expectedExpense.ID)
	if err != nil {
		t.Fatalf("GetExpense: %s", err)
	}

	if !reflect.DeepEqual(expectedExpense, gotExpense) {
		t.Fatalf("expected %v, got %v", expectedExpense, gotExpense)
	}
}

func TestUpdateExpense(t *testing.T) {
	loadFixtures()

	expenseToUpdate, err := ds.GetExpense(expenseLoyerUUID)
	if err != nil {
		t.Fatalf("GetExpense: %s", err)
	}

	expenseToUpdate.Label = "Nouveau libellé"
	expenseToUpdate.Amount = 150000
	expenseToUpdate.Date = time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC)

	err = ds.UpdateExpense(expenseToUpdate)
	if err != nil {
		t.Fatalf("UpdateExpense: %s", err)
	}

	updatedExpense, err := ds.GetExpense(expenseToUpdate.ID)
	if err != nil {
		t.Fatalf("GetExpense after update: %s", err)
	}

	if !reflect.DeepEqual(updatedExpense, expenseToUpdate) {
		t.Fatalf("UpdateExpense: expected %v, got %v", expenseToUpdate, updatedExpense)
	}
}

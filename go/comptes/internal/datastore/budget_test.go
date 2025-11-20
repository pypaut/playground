package datastore

import (
	"comptes/internal/model"
	"reflect"
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

func TestListBudgets(t *testing.T) {
	loadFixtures()

	expectedBudgets := []*model.Budget{
		{
			ID:     budgetCoursesUUID,
			Label:  "Courses",
			Amount: 45000,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
			TagID:  tagDepensesCourantesUUID,
		},
		{
			ID:     budgetEpargneChatsUUID,
			Label:  "Épargne chats",
			Amount: 4500,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
			TagID:  tagEpargneUUID,
		},
		{
			ID:     budgetCadeauUUID,
			Label:  "Cadeau pour jsp qui",
			Amount: 3900,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
			TagID:  tagDepensesVariablesUUID,
		},
		{
			ID:     budgetLoyerUUID,
			Label:  "Loyer",
			Amount: 120000,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
			TagID:  tagFacturesUUID,
		},
		{
			ID:     budgetNoelUUID,
			Label:  "Cadeaux Noël",
			Amount: 20000,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
			TagID:  tagDepensesVariablesUUID,
		},
	}

	budgets, err := ds.ListBudgets(2025, 7)
	if err != nil {
		t.Fatalf("ListBudgets: %v", err)
	}

	if len(budgets) != 5 {
		t.Fatalf("ListBudgets: got %d budgets, want 5", len(budgets))
	}

	for i := range budgets {
		if !reflect.DeepEqual(budgets[i], expectedBudgets[i]) {
			t.Fatalf("ListBudgets: got %v, want %v", budgets[i], expectedBudgets[i])
		}
	}
}

func TestListBudgetsForTag(t *testing.T) {
	loadFixtures()

	cases := []struct {
		tagId           uuid.UUID
		expectedBudgets []*model.Budget
	}{
		{
			tagId: tagEpargneUUID,
			expectedBudgets: []*model.Budget{
				{
					ID:     budgetEpargneChatsUUID,
					Label:  "Épargne chats",
					Amount: 4500,
					Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
					TagID:  tagEpargneUUID,
				},
			},
		},
		{
			tagId: tagFacturesUUID,
			expectedBudgets: []*model.Budget{
				{
					ID:     budgetLoyerUUID,
					Label:  "Loyer",
					Amount: 120000,
					Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
					TagID:  tagFacturesUUID,
				},
			},
		},
		{
			tagId: tagDepensesCourantesUUID,
			expectedBudgets: []*model.Budget{
				{
					ID:     budgetCoursesUUID,
					Label:  "Courses",
					Amount: 45000,
					Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
					TagID:  tagDepensesCourantesUUID,
				},
			},
		},
		{
			tagId: tagDepensesVariablesUUID,
			expectedBudgets: []*model.Budget{
				{
					ID:     budgetCadeauUUID,
					Label:  "Cadeau pour jsp qui",
					Amount: 3900,
					Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
					TagID:  tagDepensesVariablesUUID,
				},
				{
					ID:     budgetNoelUUID,
					Label:  "Cadeaux Noël",
					Amount: 20000,
					Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
					TagID:  tagDepensesVariablesUUID,
				},
			},
		},
	}

	for _, c := range cases {
		budgets, err := ds.ListBudgetsForTagId(c.tagId, 2025, 7)
		if err != nil {
			t.Fatalf("ListBudgets: %v", err)
		}

		if len(budgets) != len(c.expectedBudgets) {
			t.Fatalf("ListBudgets: got %d budgets, want %d", len(budgets), len(c.expectedBudgets))
		}

		for i := range budgets {
			if !reflect.DeepEqual(budgets[i], c.expectedBudgets[i]) {
				t.Fatalf("ListBudgets: got %v, want %v", budgets[i], c.expectedBudgets[i])
			}
		}

	}
}

func TestAddBudget(t *testing.T) {
	loadFixtures()

	testBudget := &model.Budget{
		Label:  "Épargne voiture",
		Amount: 150,
		Date:   time.Date(2025, 10, 9, 0, 0, 0, 0, time.UTC),
		TagID:  tagEpargneUUID,
	}

	err := ds.AddBudget(testBudget)
	if err != nil {
		t.Fatalf("AddBudget: %s", err)
	}

	budgets, err := ds.ListBudgets(2025, 10)
	if err != nil {
		t.Fatalf("ListBudgets: %v", err)
	}

	if len(budgets) != 1 {
		t.Fatalf("ListBudgets: got %d budgets, want 1", len(budgets))
	}

	if testBudget.Amount != budgets[0].Amount ||
		testBudget.Date != budgets[0].Date ||
		testBudget.TagID != budgets[0].TagID {
		t.Fatalf("ListBudgets: got %v, want %v", budgets[0], testBudget)
	}
}

func TestRemoveBudgetWithExpenses(t *testing.T) {
	loadFixtures()
	budgetId := uuid.FromStringOrNil("a853f96f-e238-49ee-97f3-1e17f0336df9")

	_, err := ds.GetBudget(budgetId)
	if err != nil {
		t.Fatalf("GetBudget: %s", err)
	}

	err = ds.RemoveBudget(budgetId)
	if err == nil {
		t.Fatalf("RemoveBudget: err should not be nil")
	}
}

func TestRemoveBudgetWithoutExpenses(t *testing.T) {
	loadFixtures()
	budgetId := uuid.FromStringOrNil("0217689a-588c-4ff0-a43a-3464f1cf69a1")

	_, err := ds.GetBudget(budgetId)
	if err != nil {
		t.Fatalf("GetBudget: %s", err)
	}

	err = ds.RemoveBudget(budgetId)
	if err != nil {
		t.Fatalf("RemoveBudget: %v", err)
	}

	budgetAfter, err := ds.GetBudget(budgetId)
	if budgetAfter != nil {
		t.Fatalf("expected nil budget, got %v", budgetAfter)
	}

	if err == nil {
		t.Fatalf("err should not be nil")
	}
}

func TestGetBudget(t *testing.T) {
	loadFixtures()

	expectedBudget := &model.Budget{
		ID:     budgetEpargneChatsUUID,
		Label:  "Épargne chats",
		Amount: 4500,
		Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
		TagID:  tagEpargneUUID,
	}

	gotBudget, err := ds.GetBudget(expectedBudget.ID)
	if err != nil {
		t.Fatalf("GetBudget: %s", err)
	}

	if !reflect.DeepEqual(expectedBudget, gotBudget) {
		t.Fatalf("expected %v, got %v", expectedBudget, gotBudget)
	}
}

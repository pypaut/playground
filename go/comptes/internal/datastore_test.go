package internal

import (
	"reflect"
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

var (
	tagDepensesCourantesUUID = uuid.FromStringOrNil("74b344cb-7a16-4af8-8b82-17f477a4f30e")
	tagEpargneUUID           = uuid.FromStringOrNil("f9383bb3-6aaf-41d7-906c-d1c580f23d49")
	tagDepensesVariablesUUID = uuid.FromStringOrNil("a4f7f30c-ae34-4480-8e28-a9ab1741dfb3")
	tagFacturesUUID          = uuid.FromStringOrNil("226cb277-5208-4a0d-8b9f-37f3630e288f")

	budgetCoursesUUID      = uuid.FromStringOrNil("a853f96f-e238-49ee-97f3-1e17f0336df9")
	budgetEpargneChatsUUID = uuid.FromStringOrNil("d253c593-440d-4bac-ac67-e4ff69355339")
	budgetCadeauUUID       = uuid.FromStringOrNil("d3d63ae4-8680-40c6-9f00-af694d83ac6d")
	budgetLoyerUUID        = uuid.FromStringOrNil("a575ca9f-ddf1-4a52-a718-c018b5169757")
)

func TestListBudgets(t *testing.T) {
	cfg, err := LoadConfig("../config.yml")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := NewDatastore(cfg)
	if err != nil {
		t.Fatal(err)
	}

	expectedBudgets := []*Budget{
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
	}

	budgets, err := ds.ListBudgets(2025, 7)
	if err != nil {
		t.Fatalf("ListBudgets: %v", err)
	}

	if len(budgets) != 4 {
		t.Fatalf("ListBudgets: got %d budgets, want 4", len(budgets))
	}

	for i := range budgets {
		if !reflect.DeepEqual(budgets[i], expectedBudgets[i]) {
			t.Fatalf("ListBudgets: got %v, want %v", budgets[i], expectedBudgets[i])
		}
	}
}

func TestListBudgetsForTag(t *testing.T) {
	cfg, err := LoadConfig("../config.yml")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := NewDatastore(cfg)
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		tag             string
		expectedBudgets []*Budget
	}{
		{
			tag: "Épargnes",
			expectedBudgets: []*Budget{
				{
					Label:  "Épargne chats",
					Amount: 4500,
					Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
					TagID:  tagEpargneUUID,
				},
			},
		},
		{
			tag: "Factures",
			expectedBudgets: []*Budget{
				{
					Label:  "Loyer",
					Amount: 120000,
					Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
					TagID:  tagFacturesUUID,
				},
			},
		},
		{
			tag: "Dépenses courantes",
			expectedBudgets: []*Budget{
				{
					Label:  "Courses",
					Amount: 45000,
					Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
					TagID:  tagDepensesCourantesUUID,
				},
			},
		},
		{
			tag: "Dépenses variables",
			expectedBudgets: []*Budget{
				{
					Label:  "Cadeau pour jsp qui",
					Amount: 3900,
					Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
					TagID:  tagDepensesCourantesUUID,
				},
			},
		},
	}

	for _, c := range cases {
		budgets, err := ds.ListBudgetsForTag(c.tag, 2025, 7)
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

func TestListIncomes(t *testing.T) {
	cfg, err := LoadConfig("../config.yml")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := NewDatastore(cfg)
	if err != nil {
		t.Fatal(err)
	}

	expectedIncomes := []*Income{
		{
			Label:  "Salaire 1",
			Amount: 200042,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Label:  "Salaire 2",
			Amount: 210081,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	incomes, err := ds.ListIncomes(2025, 7)
	if err != nil {
		t.Fatalf("ListIncomes: %v", err)
	}

	if len(incomes) != 2 {
		t.Fatalf("ListIncomes: got %d incomes, want 2", len(incomes))
	}

	for i := range incomes {
		if !reflect.DeepEqual(incomes[i], expectedIncomes[i]) {
			t.Fatalf("ListIncomes: got %v, want %v", incomes[i], expectedIncomes[i])
		}
	}
}

func TestListTags(t *testing.T) {
	cfg, err := LoadConfig("../config.yml")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := NewDatastore(cfg)
	if err != nil {
		t.Fatal(err)
	}

	expectedTags := []*Tag{
		{
			Label:       "Factures",
			Description: "Paiements récurrents, charges fixes, abonnements",
			Icon:        "🧾",
		},
		{
			Label:       "Épargnes",
			Description: "On met de côté",
			Icon:        "💰",
		},
		{
			Label:       "Dépenses courantes",
			Description: "Dépenses usuelles",
			Icon:        "💳",
		},
		{
			Label:       "Dépenses variables",
			Description: "Dépenses variables",
			Icon:        "💶",
		},
	}

	tags, err := ds.ListTags()
	if err != nil {
		t.Fatalf("ListTags: %v", err)
	}

	if len(tags) != 4 {
		t.Fatalf("ListTags: got %d tags, want 4", len(tags))
	}

	for i := range tags {
		if !reflect.DeepEqual(expectedTags[i], tags[i]) {
			t.Fatalf("ListTags: got %v, want %v", tags[i], expectedTags[i])
		}
	}
}

func TestListExpenses(t *testing.T) {
	cfg, err := LoadConfig("../config.yml")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := NewDatastore(cfg)
	if err != nil {
		t.Fatal(err)
	}

	budgetCoursesUUID := uuid.FromStringOrNil("a853f96f-e238-49ee-97f3-1e17f0336df9")
	budgetLoyerUUID := uuid.FromStringOrNil("a575ca9f-ddf1-4a52-a718-c018b5169757")

	expectedExpenses := []*Expense{
		{
			Label:    "Loyer",
			Amount:   120000,
			Date:     time.Date(2025, 07, 2, 0, 0, 0, 0, time.UTC),
			BudgetID: budgetLoyerUUID,
		},
		{
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

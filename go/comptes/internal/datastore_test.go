package internal

import (
	"reflect"
	"testing"
	"time"
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
			Label:  "Courses",
			Amount: 45000,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
			Tag:    "Dépenses courantes",
		},
		{
			Label:  "Épargne chats",
			Amount: 4500,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
			Tag:    "Épargnes",
		},
		{
			Label:  "Cadeau pour jsp qui",
			Amount: 3900,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
			Tag:    "Dépenses variables",
		},
		{
			Label:  "Loyer",
			Amount: 120000,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
			Tag:    "Factures",
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
					Tag:    "Épargnes",
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
					Tag:    "Factures",
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
					Tag:    "Dépenses courantes",
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
					Tag:    "Dépenses variables",
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

	expectedExpenses := []*Expense{
		{
			Label:  "Loyer",
			Amount: 120000,
			Date:   time.Date(2025, 07, 2, 0, 0, 0, 0, time.UTC),
			Budget: "Loyer",
		},
		{
			Label:  "Leclerc",
			Amount: 4781,
			Date:   time.Date(2025, 07, 8, 0, 0, 0, 0, time.UTC),
			Budget: "Courses",
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

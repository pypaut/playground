package internal

import "testing"

func TestListBudgets(t *testing.T) {
	cfg, err := LoadConfig("../config.yml")
	if err != nil {
		t.Fatal(err)
	}

	ds, err := NewDatastore(cfg)
	if err != nil {
		t.Fatal(err)
	}

	budgets, err := ds.ListBudgets()
	if err != nil {
		t.Fatalf("ListBudgets: %v", err)
	}

	if len(budgets) != 4 {
		t.Fatalf("ListBudgets: got %d budgets, want 4", len(budgets))
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

	incomes, err := ds.ListIncomes()
	if err != nil {
		t.Fatalf("ListIncomes: %v", err)
	}

	if len(incomes) != 2 {
		t.Fatalf("ListIncomes: got %d incomes, want 2", len(incomes))
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

	tags, err := ds.ListTags()
	if err != nil {
		t.Fatalf("ListTags: %v", err)
	}

	if len(tags) != 4 {
		t.Fatalf("ListTags: got %d tags, want 4", len(tags))
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

	expenses, err := ds.ListExpenses()
	if err != nil {
		t.Fatalf("ListExpenses: %v", err)
	}

	if len(expenses) != 2 {
		t.Fatalf("ListExpenses: got %d expenses, want 2", len(expenses))
	}
}

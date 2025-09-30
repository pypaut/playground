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

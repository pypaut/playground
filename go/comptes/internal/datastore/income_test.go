package datastore

import (
	"comptes/internal/model"
	"reflect"
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

func TestListIncomes(t *testing.T) {
	loadFixtures()

	expectedIncomes := []*model.Income{
		{
			ID:     uuid.FromStringOrNil("961a1dd1-ca6f-412a-83f0-9af6dcd85081"),
			Label:  "Salaire 1",
			Amount: 200042,
			Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:     uuid.FromStringOrNil("247a13b0-32bc-4dd2-8250-b9beabfc939f"),
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

func TestGetIncome(t *testing.T) {
	loadFixtures()

	expectedIncome := &model.Income{
		ID:     uuid.FromStringOrNil("961a1dd1-ca6f-412a-83f0-9af6dcd85081"),
		Label:  "Salaire 1",
		Amount: 200042,
		Date:   time.Date(2025, 07, 1, 0, 0, 0, 0, time.UTC),
	}

	gotIncome, err := ds.GetIncome(expectedIncome.ID)
	if err != nil {
		t.Fatalf("GetIncome: %s", err)
	}

	if !reflect.DeepEqual(expectedIncome, gotIncome) {
		t.Fatalf("expected %v, got %v", expectedIncome, gotIncome)
	}
}

func TestAddIncome(t *testing.T) {
	loadFixtures()

	testIncome := &model.Income{
		Label:  "Salaire 3",
		Amount: 250000,
		Date:   time.Date(2025, 10, 9, 0, 0, 0, 0, time.UTC),
	}

	err := ds.AddIncome(testIncome)
	if err != nil {
		t.Fatalf("AddIncome: %s", err)
	}

	incomes, err := ds.ListIncomes(2025, 10)
	if err != nil {
		t.Fatalf("ListIncomes: %v", err)
	}

	if len(incomes) != 1 {
		t.Fatalf("ListIncomes: got %d incomes, want 1", len(incomes))
	}

	if testIncome.Amount != incomes[0].Amount ||
		testIncome.Date != incomes[0].Date {
		t.Fatalf("ListIncomes: got %v, want %v", incomes[0], testIncome)
	}
}

func TestRemoveIncome(t *testing.T) {
	loadFixtures()

	incomeId := uuid.FromStringOrNil("961a1dd1-ca6f-412a-83f0-9af6dcd85081")

	_, err := ds.GetIncome(incomeId)
	if err != nil {
		t.Fatalf("GetIncome: %s", err)
	}

	err = ds.RemoveIncome(incomeId)
	if err != nil {
		t.Fatalf("RemoveIncome: %v", err)
	}

	incomeAfter, err := ds.GetIncome(incomeId)
	if incomeAfter != nil {
		t.Fatalf("expected nil income, got %v", incomeAfter)
	}

	if err == nil {
		t.Fatalf("err should not be nil")
	}
}

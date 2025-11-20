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


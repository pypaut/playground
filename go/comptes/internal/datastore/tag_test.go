package datastore

import (
	"comptes/internal/model"
	"reflect"
	"testing"

	"github.com/gofrs/uuid"
)

func TestListTags(t *testing.T) {
	loadFixtures()

	expectedTags := []*model.Tag{
		{
			ID:          tagFacturesUUID,
			Label:       "Factures",
			Description: "Paiements récurrents, charges fixes, abonnements",
			Icon:        "🧾",
		},
		{
			ID:          tagEpargneUUID,
			Label:       "Épargnes",
			Description: "On met de côté",
			Icon:        "💰",
		},
		{
			ID:          tagDepensesCourantesUUID,
			Label:       "Dépenses courantes",
			Description: "Dépenses usuelles",
			Icon:        "💳",
		},
		{
			ID:          tagDepensesVariablesUUID,
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

func TestGetTagByLabel(t *testing.T) {
	loadFixtures()

	cases := []struct {
		label      string
		expectedId uuid.UUID
	}{
		{
			label:      "Factures",
			expectedId: tagFacturesUUID,
		},
		{
			label:      "Épargnes",
			expectedId: tagEpargneUUID,
		},
		{
			label:      "Dépenses courantes",
			expectedId: tagDepensesCourantesUUID,
		},
		{
			label:      "Dépenses variables",
			expectedId: tagDepensesVariablesUUID,
		},
	}

	for _, c := range cases {
		gotTag, err := ds.GetTagByLabel(c.label)
		if err != nil {
			t.Fatalf("GetTagByLabel error: %s", err)
		}

		if gotTag.ID != c.expectedId {
			t.Fatalf("GetTagByLabel: got %v, want %v", gotTag.ID, c.expectedId)
		}
	}
}

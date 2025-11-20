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

func TestGetTag(t *testing.T) {
	loadFixtures()

	expectedTag := &model.Tag{
		ID:          tagFacturesUUID,
		Label:       "Factures",
		Description: "Paiements récurrents, charges fixes, abonnements",
		Icon:        "🧾",
	}

	gotTag, err := ds.GetTag(expectedTag.ID)
	if err != nil {
		t.Fatalf("GetTag: %s", err)
	}

	if !reflect.DeepEqual(expectedTag, gotTag) {
		t.Fatalf("expected %v, got %v", expectedTag, gotTag)
	}
}

func TestAddTag(t *testing.T) {
	loadFixtures()

	testTag := &model.Tag{
		Label:       "Test",
		Description: "Tag de test",
		Icon:        "🧪",
	}

	err := ds.AddTag(testTag)
	if err != nil {
		t.Fatalf("AddTag: %s", err)
	}

	tags, err := ds.ListTags()
	if err != nil {
		t.Fatalf("ListTags: %v", err)
	}

	// Trouver le tag ajouté
	var addedTag *model.Tag
	for _, tag := range tags {
		if tag.Label == "Test" {
			addedTag = tag
			break
		}
	}

	if addedTag == nil {
		t.Fatalf("AddTag: tag not found in list")
	}

	if testTag.Label != addedTag.Label ||
		testTag.Description != addedTag.Description ||
		testTag.Icon != addedTag.Icon {
		t.Fatalf("AddTag: got %v, want %v", addedTag, testTag)
	}
}

func TestRemoveTag(t *testing.T) {
	loadFixtures()

	tagId := tagFacturesUUID

	_, err := ds.GetTag(tagId)
	if err != nil {
		t.Fatalf("GetTag: %s", err)
	}

	err = ds.RemoveTag(tagId)
	if err == nil {
		t.Fatal("RemoveTag: err should not be nil")
	}
}

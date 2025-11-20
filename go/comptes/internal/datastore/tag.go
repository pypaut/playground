package datastore

import (
	"comptes/internal/model"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

func (d *Datastore) GetTagByLabel(tagLabel string) (*model.Tag, error) {
	var tag model.Tag
	err := d.dbpool.QueryRow(context.Background(), "select * from tags where label = $1", tagLabel).Scan(
		&tag.ID, &tag.Label, &tag.Description, &tag.Icon,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get tag: %w", err)
	}

	return &tag, nil
}

func (d *Datastore) ListTags() (tags []*model.Tag, err error) {
	rows, err := d.dbpool.Query(context.Background(), "select * from tags")
	if err != nil {
		return nil, fmt.Errorf("could not list tags: %w", err)
	}

	for rows.Next() {
		var tag model.Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Label,
			&tag.Description,
			&tag.Icon,
		)

		if err != nil {
			return nil, fmt.Errorf("could not scan tags: %w", err)
		}

		tags = append(tags, &tag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate tags: %w", err)
	}

	return
}

func (d *Datastore) GetTag(tagId uuid.UUID) (*model.Tag, error) {
	var tag model.Tag
	err := d.dbpool.QueryRow(context.Background(), "select * from tags where id = $1", tagId).Scan(
		&tag.ID,
		&tag.Label,
		&tag.Description,
		&tag.Icon,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get tag: %w", err)
	}

	return &tag, nil
}

func (d *Datastore) AddTag(tag *model.Tag) error {
	query := `INSERT INTO tags (label, description, icon)
VALUES (@label, @description, @icon)`
	args := pgx.NamedArgs{
		"label":       tag.Label,
		"description": tag.Description,
		"icon":        tag.Icon,
	}

	_, err := d.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("error creating tag: %s", err)
	}

	return nil
}

func (d *Datastore) RemoveTag(tagId uuid.UUID) error {
	query := `DELETE FROM tags WHERE id=@tagId`
	args := pgx.NamedArgs{
		"tagId": tagId,
	}

	_, err := d.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("error removing tag: %s", err)
	}

	return nil
}

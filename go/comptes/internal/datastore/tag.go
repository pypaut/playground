package datastore

import (
	"comptes/internal/model"
	"context"
	"fmt"
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

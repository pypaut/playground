package cli

import (
	"comptes/internal/datastore"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type GooseAction string

const (
	GooseActionUp     GooseAction = "up"
	GooseActionDown   GooseAction = "down"
	GooseActionStatus GooseAction = "status"
)

type GooseCmd struct {
	datastore *datastore.Datastore

	Action GooseAction
}

func (cmd *GooseCmd) Run() error {
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	sqlDB := stdlib.OpenDBFromPool(cmd.datastore.GetPool())
	defer sqlDB.Close()

	switch cmd.Action {
	case GooseActionUp:
		if err := goose.Up(sqlDB, "./migrations"); err != nil {
			panic(err)
		}

	case GooseActionDown:
		if err := goose.Down(sqlDB, "./migrations"); err != nil {
			panic(err)
		}

	case GooseActionStatus:
		if err := goose.Status(sqlDB, "./migrations"); err != nil {
			panic(err)
		}
	}

	return nil
}

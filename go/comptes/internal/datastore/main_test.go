package datastore

import (
	"comptes/internal/config"
	"database/sql"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
)

var (
	ds       *Datastore
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	cfg, err := config.LoadConfig("../../config.yml")
	if err != nil {
		panic(err)
	}

	ds, err = NewDatastore(cfg)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", cfg.Db.Connection)
	if err != nil {
		panic(err)
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../../fixtures"),
	)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func loadFixtures() {
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}

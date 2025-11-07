package main

import (
	"comptes/internal/cli"
	"comptes/internal/config"
	"comptes/internal/datastore"

	"github.com/alecthomas/kong"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		panic(err)
	}

	ds, err := datastore.NewDatastore(cfg)
	if err != nil {
		panic(err)
	}

	c := cli.NewCli(ds)
	ctx := kong.Parse(c)
	err = ctx.Run(c)
	ctx.FatalIfErrorf(err)

}

package main

import (
	"comptes/internal"

	"github.com/alecthomas/kong"
)

func main() {
	cfg, err := internal.LoadConfig("config.yml")
	if err != nil {
		panic(err)
	}

	datastore, err := internal.NewDatastore(cfg)
	if err != nil {
		panic(err)
	}

	cli := internal.NewCli(datastore)
	ctx := kong.Parse(cli)
	err = ctx.Run(cli)
	ctx.FatalIfErrorf(err)

}

package main

import (
	"github.com/ngocthanh1389/orders/be/pkg/postgres"
	"github.com/urfave/cli"
)

const (
	// server
	bindAddressFlag    = "bind-address"
	defaultBindAddress = "0.0.0.0:1234"
)

func NewAppFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   bindAddressFlag,
			Usage:  "provide application host and port",
			EnvVar: "BIND_ADDRESS",
			Value:  defaultBindAddress,
		},
	}
}

func NewApp() *cli.App {
	app := cli.NewApp()

	app.Flags = NewAppFlags()
	app.Flags = append(app.Flags, postgres.PostgresSQLFlags()...)

	return app
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ngocthanh1389/orders/be/pkg/postgres"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	app := NewApp()
	app.Name = "X-Monitor"
	app.Version = "1.1.0"
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}

func run(c *cli.Context) error {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	// database
	db, err := postgres.NewDBFromContext(c)
	if err != nil {
		log.Panic("cannot init DB connection", "err", err)
	}
	_, err = postgres.RunMigrationUp(
		db.DB, c.String(postgres.PostgresMigrationPath), c.String(postgres.PostgresDatabaseFlag))
	if err != nil {
		log.Panic("cannot init DB", "err", err)
	}

	// http client
	// httpClient := &http.Client{
	// 	Timeout: time.Second * 30,
	// 	Transport: &http.Transport{
	// 		IdleConnTimeout:       time.Second * 120,
	// 		ResponseHeaderTimeout: time.Second * 10,
	// 	},
	// }
	return nil
}

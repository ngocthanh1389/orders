package main

import (
	"log"
	"os"

	"github.com/ngocthanh1389/orders/be/internal/service"
	"github.com/ngocthanh1389/orders/be/pkg/postgres"
	"github.com/ngocthanh1389/orders/be/pkg/server"
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
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"stdout",
		"app.log",
	}
	rl, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	l := rl.Sugar()

	l.Infow("Starting app...")
	// database
	db, err := postgres.NewDBFromContext(c)
	if err != nil {
		l.Panicw("cannot init DB connection", "error", err)
	}
	if _, err := postgres.RunMigrationUp(
		db.DB, c.String(postgres.PostgresMigrationPath), c.String(postgres.PostgresDatabaseFlag)); err != nil {
		l.Panicw("cannot migrate DB", "error", err)
	}

	// http client
	// httpClient := &http.Client{
	// 	Timeout: time.Second * 30,
	// 	Transport: &http.Transport{
	// 		IdleConnTimeout:       time.Second * 120,
	// 		ResponseHeaderTimeout: time.Second * 10,
	// 	},
	// }
	server := server.NewServer(c.String(bindAddressFlag))
	service.AddOrderService(server, db, l)
	return server.Run()
}

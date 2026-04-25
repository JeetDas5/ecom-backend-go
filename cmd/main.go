package main

import (
	"context"
	"github/JeetDas5/ecom-app/internal/env"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString(
				"GOOSE_DBSTRING",
				"host=localhost port=5434 user=postgres password=postgres dbname=ecom sslmode=disable",
			),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	//print db string
	logger.Info("database connection string", "dsn", cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db: conn,
	}

	slog.Info("starting server", "addr", api.config.addr)

	if err := api.run(api.mount()); err != nil {
		log.Printf("server failed to start: %s", err)
		os.Exit(1)
	}
}

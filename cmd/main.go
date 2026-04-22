package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("starting server", "addr", api.config.addr)

	if err := api.run(api.mount()); err != nil {
		log.Printf("server failed to start: %s", err)
		os.Exit(1)
	}
}

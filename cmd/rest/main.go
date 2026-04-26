package main

import (
	"context"
	"log/slog"

	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/rest"
)

func main() {
	slog.Info("Starting rest server")

	slog.Debug("Getting config")
	c := config.Get(true)

	slog.Debug("Migrating")
	ctx := context.Background()
	err := db.Migrate(ctx, c)
	if err != nil {
		panic(err)
	}

	r := rest.NewRest()
	err = r.Start()
	if err != nil {
		panic(err)
	}
}

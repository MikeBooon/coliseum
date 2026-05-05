package main

import (
	"context"
	"log/slog"

	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/rest"
	"github.com/MikeBooon/coliseum/internal/system"
	"github.com/MikeBooon/coliseum/service"
)

func main() {
	slog.Info("Starting rest server")

	slog.Debug("Getting config")
	env := config.GetEnv(true)

	slog.Debug("Migrating")
	ctx := context.Background()
	err := db.Migrate(ctx, env)
	if err != nil {
		panic(err)
	}

	db := db.Connect(env)

	sys := system.System{
		DB:        db,
		EnvConfig: env,
		Svcs:      service.NewServices(db),
	}
	r := rest.NewRest(sys)
	err = r.Start()
	if err != nil {
		panic(err)
	}
}

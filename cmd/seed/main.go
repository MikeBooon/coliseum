package main

import (
	"context"
	"log"
	"log/slog"
	"maps"
	"os"
	"slices"

	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/system"
)

var seedFiles = map[string]func(sys system.System) error{}

func main() {
	slog.Info("Seed starting")

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
	}

	if len(os.Args) <= 1 {
		log.Fatalf("Seed file not specified")
	}
	seed := os.Args[1]

	if !slices.Contains(slices.Sorted(maps.Keys(seedFiles)), seed) {
		log.Fatalf("Seed file '%s' not found", seed)
	}

	seedFunc := seedFiles[seed]

	err = seedFunc(sys)

	if err != nil {
		log.Fatalf("Seed file failed: %v", err)
	}

	println(seed)
}

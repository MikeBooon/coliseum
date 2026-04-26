package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/system"
)

type IntegrationTestDeps struct {
	System      system.System
	pgContainer *PostgresContainer
}

func RunTestWithIntegrationDeps(
	c **IntegrationTestDeps,
	m *testing.M,
) {
	ctx := context.Background()
	*c = initIntegrationTestDeps(ctx)

	exitCode := m.Run()
	(*c).Teardown(ctx)
	os.Exit(exitCode)
}

func initIntegrationTestDeps(ctx context.Context) *IntegrationTestDeps {
	pg, err := createTestPGContainer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	c := &config.Config{
		DBConn: pg.ConnectionString,
	}
	if err = db.Migrate(ctx, c); err != nil {
		log.Fatal(err)
	}
	db := db.Connect(c)
	sys := system.System{
		DB:     db,
		Config: c,
	}

	return &IntegrationTestDeps{
		System:      sys,
		pgContainer: pg,
	}
}

func (deps *IntegrationTestDeps) Teardown(ctx context.Context) {
	destroyTestPGContainer(ctx, deps.pgContainer)
}

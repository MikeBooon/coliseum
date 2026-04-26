package test

import (
	"context"

	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func createTestPGContainer(ctx context.Context) (*PostgresContainer, error) {
	pgContainer, err := postgres.Run(ctx,
		"postgres:18-alpine",
		postgres.WithDatabase("lavender-test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		return nil, err
	}
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}, nil
}

func destroyTestPGContainer(ctx context.Context, container *PostgresContainer) {
	container.Terminate(ctx)
}

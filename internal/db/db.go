// Package db contains database interactions
package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/MikeBooon/coliseum/internal/db/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

const MAX_CONNECTIONS = 20
const MAX_IDLE_CONNECTIONS = 40
const MAX_CONNECTION_LIFETIME = 3 * time.Minute
const MAX_CONNECTION_IDLE_TIME = 3 * time.Minute

type DB = bun.DB
type IDB = bun.IDB

func Migrate(ctx context.Context, c *config.Config) error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(c.DBConn),
	))
	db := bun.NewDB(sqldb, pgdialect.New())

	migrator := migrate.NewMigrator(db, migrations.Migrations)
	if err := migrator.Init(ctx); err != nil {
		return err
	}

	_, err := migrator.Migrate(ctx)

	return err
}

func Connect(c *config.Config) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(c.DBConn),
	))
	sqldb.SetMaxOpenConns(MAX_CONNECTIONS)
	sqldb.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
	sqldb.SetConnMaxLifetime(MAX_CONNECTION_LIFETIME)
	sqldb.SetConnMaxIdleTime(MAX_CONNECTION_IDLE_TIME)

	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

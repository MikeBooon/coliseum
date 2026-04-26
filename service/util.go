package service

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func isNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func whereTenant(tenantID uuid.UUID) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("tenant_id = ?", tenantID)
	}
}

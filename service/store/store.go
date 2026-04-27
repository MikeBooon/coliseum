package store

import (
	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TenantStore struct {
	db       db.IDB
	tenantID uuid.UUID
}

func (store TenantStore) NewSelect(
	model any,
) *bun.SelectQuery {
	return store.db.NewSelect().
		Model(model).
		Apply(whereTenant(store.tenantID))
}

func (store TenantStore) NewInsert(
	model any,
) *bun.InsertQuery {
	if tm, ok := model.(dao.TenantSetter); ok {
		tm.SetTenantID(store.tenantID)
	}
	return store.db.NewInsert().
		Model(model)
}

func whereTenant(tenantID uuid.UUID) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("tenant_id = ?", tenantID)
	}
}

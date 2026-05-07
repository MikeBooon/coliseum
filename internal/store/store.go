package store

import (
	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TenantStore struct {
	db       db.IDB
	TenantID uuid.UUID
}

func (store TenantStore) NewSelect(
	model any,
) *bun.SelectQuery {
	sel := store.db.NewSelect().
		Model(model).
		Apply(whereTenant(store.TenantID))

	if _, ok := model.(dao.TenantSetter); ok {
		sel.Apply(whereTenant(store.TenantID))
	}

	return sel
}

func (store TenantStore) NewInsert(
	model any,
) *bun.InsertQuery {
	if tm, ok := model.(dao.TenantSetter); ok {
		tm.SetTenantID(store.TenantID)
	}
	return store.db.NewInsert().
		Model(model)
}

func (store TenantStore) NewUpdate(
	model any,
) *bun.UpdateQuery {
	update := store.db.NewUpdate().
		Model(model)

	if _, ok := model.(dao.TenantSetter); ok {
		update.Apply(whereTenantUpdate(store.TenantID))
	}

	return update
}

func (store TenantStore) NewDelete(
	model any,
) *bun.DeleteQuery {
	del := store.db.NewDelete().
		Model(model)

	if _, ok := model.(dao.TenantSetter); ok {
		del.Apply(whereTenantDelete(store.TenantID))
	}

	return del
}

func whereTenant(tenantID uuid.UUID) func(*bun.SelectQuery) *bun.SelectQuery {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("tenant_id = ?", tenantID)
	}
}

func whereTenantUpdate(tenantID uuid.UUID) func(*bun.UpdateQuery) *bun.UpdateQuery {
	return func(q *bun.UpdateQuery) *bun.UpdateQuery {
		return q.Where("tenant_id = ?", tenantID)
	}
}

func whereTenantDelete(tenantID uuid.UUID) func(*bun.DeleteQuery) *bun.DeleteQuery {
	return func(q *bun.DeleteQuery) *bun.DeleteQuery {
		return q.Where("tenant_id = ?", tenantID)
	}
}

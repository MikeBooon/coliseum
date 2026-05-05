package store

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/tenant"
)

type Provider struct {
	db db.IDB
}

func NewProvider(db db.IDB) *Provider {
	return &Provider{
		db: db,
	}
}

func (sp Provider) NewTenantStore(ctx context.Context) TenantStore {
	return TenantStore{
		db:       sp.db,
		tenantID: tenant.MustFromContext(ctx),
	}
}

func (sp Provider) DB() db.IDB {
	return sp.db
}

package service

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/service/store"
	"github.com/uptrace/bun"
)

type Services struct {
	db     db.IDB
	User   *UserService
	Tenant *TenantService
	RBAC   *RBACService
}

func NewServices(db db.IDB) *Services {
	storeProvider := store.NewProvider(db)

	return &Services{
		db:     db,
		User:   &UserService{storeProvider},
		Tenant: &TenantService{storeProvider},
		RBAC:   &RBACService{storeProvider},
	}
}

func (s *Services) InTx(ctx context.Context, fn func(s *Services) error) error {
	return s.db.(*bun.DB).RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return fn(NewServices(tx))
	})
}

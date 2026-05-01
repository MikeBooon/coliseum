package service

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/service/store"
	"github.com/uptrace/bun"
)

type Services struct {
	db        db.IDB
	User      *UserService
	Tenant    *TenantService
	RBAC      *RBACService
	Provision *ProvisionService
}

func NewServices(db db.IDB) *Services {
	storeProvider := store.NewProvider(db)

	svcs := &Services{}

	svcs.db = db
	svcs.RBAC = &RBACService{storeProvider}
	svcs.User = &UserService{storeProvider}
	svcs.Tenant = &TenantService{storeProvider}
	svcs.Provision = &ProvisionService{storeProvider, svcs}

	return svcs
}

func (s *Services) InTx(ctx context.Context, fn func(s *Services) error) error {
	return s.db.(*bun.DB).RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return fn(NewServices(tx))
	})
}

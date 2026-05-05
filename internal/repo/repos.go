package repo

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/store"
	"github.com/uptrace/bun"
)

type Repos struct {
	db     db.IDB
	User   *UserRepo
	Tenant *TenantRepo
	RBAC   *RBACRepo
}

func NewRepos(db db.IDB) *Repos {
	storeProvider := store.NewProvider(db)

	svcs := &Repos{}

	svcs.db = db
	svcs.RBAC = &RBACRepo{storeProvider}
	svcs.User = &UserRepo{storeProvider}
	svcs.Tenant = &TenantRepo{storeProvider}

	return svcs
}

func (s *Repos) InTx(ctx context.Context, fn func(s *Repos) error) error {
	return s.db.(*bun.DB).RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return fn(NewRepos(tx))
	})
}

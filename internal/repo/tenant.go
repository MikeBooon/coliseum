package repo

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/MikeBooon/coliseum/internal/store"
)

type TenantRepo struct {
	*store.Provider
}

func (r TenantRepo) New(ctx context.Context, name string) (*dao.Tenant, error) {
	t := &dao.Tenant{Name: name}
	err := r.DB().NewInsert().Model(t).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

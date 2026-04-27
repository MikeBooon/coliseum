package service

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/MikeBooon/coliseum/service/store"
)

type TenantService struct {
	*store.Provider
}

func (s TenantService) New(ctx context.Context, name string) (*dao.Tenant, error) {
	t := &dao.Tenant{Name: name}
	err := s.DB().NewInsert().Model(t).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

package service

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/db/dao"
)

type TenantService struct {
	db db.IDB
}

func (s TenantService) New(ctx context.Context, name string) (*dao.Tenant, error) {
	t := &dao.Tenant{Name: name}
	err := s.db.NewInsert().Model(t).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

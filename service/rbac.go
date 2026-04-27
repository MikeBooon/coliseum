package service

import (
	"context"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/MikeBooon/coliseum/service/store"
	"github.com/google/uuid"
)

type RBACService struct {
	*store.Provider
}

func (s RBACService) NewRole(
	ctx context.Context,
	name string,
	userType domain.UserType,
) (*dao.Role, error) {
	tStore := s.NewTenantStore(ctx)
	role := &dao.Role{Name: name, Type: userType}
	err := tStore.NewInsert(role).Scan(ctx)
	return role, err
}

func (s RBACService) GetRole(
	ctx context.Context,
	id uuid.UUID,
) (*dao.Role, error) {
	tStore := s.NewTenantStore(ctx)
	role := new(dao.Role)
	err := tStore.NewSelect(role).Where("id = ?", id).Scan(ctx)
	return role, err
}

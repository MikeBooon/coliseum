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

type NewRoleOpts struct {
	Name      string
	UserType  domain.UserType
	IsDefault bool
}

func (s RBACService) NewRole(
	ctx context.Context,
	opts NewRoleOpts,
) (*dao.Role, error) {
	tStore := s.NewTenantStore(ctx)
	role := &dao.Role{Name: opts.Name, Type: opts.UserType, Default: opts.IsDefault}
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

func (s RBACService) SetUserRole(
	ctx context.Context,
	userID uuid.UUID,
	roleID uuid.UUID,
) (*dao.User, error) {
	user := new(dao.User)
	err := s.DB().NewUpdate().
		Model(user).
		Where("id = ?", userID).
		Set("role_id = ?", roleID).
		Scan(ctx)
	return user, err
}

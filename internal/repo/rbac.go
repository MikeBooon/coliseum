package repo

import (
	"context"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/MikeBooon/coliseum/internal/store"
	"github.com/google/uuid"
)

type RBACRepo struct {
	*store.Provider
}

type NewRoleOpts struct {
	Name      string
	UserType  domain.UserType
	IsDefault bool
}

func (r RBACRepo) NewRole(
	ctx context.Context,
	opts NewRoleOpts,
) (*dao.Role, error) {
	tStore := r.NewTenantStore(ctx)
	role := &dao.Role{Name: opts.Name, Type: opts.UserType, Default: opts.IsDefault}
	err := tStore.NewInsert(role).Scan(ctx)
	return role, err
}

func (r RBACRepo) GetRole(
	ctx context.Context,
	id uuid.UUID,
) (*dao.Role, error) {
	tStore := r.NewTenantStore(ctx)
	role := new(dao.Role)
	err := tStore.NewSelect(role).Where("id = ?", id).Scan(ctx)
	return role, err
}

func (r RBACRepo) SetUserRole(
	ctx context.Context,
	userID uuid.UUID,
	roleID uuid.UUID,
) (*dao.User, error) {
	tStore := r.NewTenantStore(ctx)
	user := new(dao.User)
	err := tStore.NewUpdate(user).
		Where("id = ?", userID).
		Set("role_id = ?", roleID).
		Scan(ctx)
	return user, err
}

func (r RBACRepo) ClearRolePermissions(
	ctx context.Context,
	roleID uuid.UUID,
) error {
	tStore := r.NewTenantStore(ctx)

	err := tStore.NewSelect(&dao.Role{}).
		Where("role_id = ?", roleID).
		Scan(ctx)

	if isNoRows(err) {
		return ErrRoleNotFound
	}
	if err != nil {
		return err
	}

	err = tStore.NewDelete(&dao.Permission{}).
		Where("role_id = ?", roleID).
		Scan(ctx)

	if isNoRows(err) {
		return nil
	}

	return err
}

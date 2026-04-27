package service

import (
	"context"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/MikeBooon/coliseum/internal/tenant"
	"github.com/google/uuid"
)

type RBACService struct {
	db db.IDB
}

func (s RBACService) NewRole(
	ctx context.Context,
	name string,
	userType domain.UserType,
) (*dao.Role, error) {
	tID := tenant.MustFromContext(ctx)
	tStore := newTenantStore(s.db, tID)
	role := &dao.Role{Name: name, Type: userType}
	tStore.newInsert(role)
	err := s.db.NewInsert().Model(role).Scan(ctx)
	return role, err
}

func (s RBACService) GetRole(
	ctx context.Context,
	id uuid.UUID,
) (*dao.Role, error) {
	tID := tenant.MustFromContext(ctx)
	tenantStore := newTenantStore(s.db, tID)
	role := new(dao.Role)
	err := tenantStore.newSelect(role).Where("id = ?", id).Scan(ctx)
	return role, err
}

package service

import (
	"context"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/MikeBooon/coliseum/internal/tenant"
	"github.com/MikeBooon/coliseum/service/store"
)

type ProvisionService struct {
	*store.Provider
	svcs *Services
}

type ScaffoldTenantOpts struct {
	TenantName string
	UserEmail  string
}

func (s *ProvisionService) ScaffoldTenant(
	ctx context.Context,
	opts ScaffoldTenantOpts,
) (*dao.Tenant, error) {
	t := new(dao.Tenant)
	err := s.svcs.InTx(ctx, func(svcs *Services) error {
		t, err := s.svcs.Tenant.New(ctx, opts.TenantName)
		if err != nil {
			return err
		}

		tCtx := tenant.NewContext(ctx, t.ID)

		adminRole, err := s.svcs.RBAC.NewRole(tCtx, NewRoleOpts{
			Name:      "Admin",
			UserType:  domain.TenantUserType,
			IsDefault: true,
		})
		if err != nil {
			return err
		}

		_, err = s.svcs.User.New(tCtx, NewUserOpts{
			Email:  opts.UserEmail,
			Type:   domain.TenantUserType,
			RoleID: adminRole.ID,
		})
		if err != nil {
			return err
		}

		_, err = s.svcs.RBAC.NewRole(tCtx, NewRoleOpts{
			Name:      "Admin",
			UserType:  domain.ClientUserType,
			IsDefault: true,
		})

		return err
	})

	return t, err
}

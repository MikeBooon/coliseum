package service

import (
	"context"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/MikeBooon/coliseum/internal/repo"
	"github.com/MikeBooon/coliseum/internal/tenant"
)

type ProvisionService struct {
	repos *repo.Repos
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
	err := s.repos.InTx(ctx, func(repos *repo.Repos) error {
		t, err := repos.Tenant.New(ctx, opts.TenantName)
		if err != nil {
			return err
		}

		tCtx := tenant.NewContext(ctx, t.ID)

		adminRole, err := repos.RBAC.NewRole(tCtx, repo.NewRoleOpts{
			Name:      "Admin",
			UserType:  domain.TenantUserType,
			IsDefault: true,
		})
		if err != nil {
			return err
		}

		_, err = repos.User.New(tCtx, repo.NewUserOpts{
			Email:  opts.UserEmail,
			Type:   domain.TenantUserType,
			RoleID: adminRole.ID,
		})
		if err != nil {
			return err
		}

		_, err = repos.RBAC.NewRole(tCtx, repo.NewRoleOpts{
			Name:      "Admin",
			UserType:  domain.ClientUserType,
			IsDefault: true,
		})

		return err
	})

	return t, err
}

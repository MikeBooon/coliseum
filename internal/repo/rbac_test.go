package repo_test

import (
	"testing"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/MikeBooon/coliseum/internal/repo"
	"github.com/MikeBooon/coliseum/internal/tenant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRole(t *testing.T) {
	ctx := t.Context()
	r := testDeps.System.Repos

	tnt, err := r.Tenant.New(ctx, "TestNewRole")

	require.NoError(t, err)

	tCtx := tenant.NewContext(ctx, tnt.ID)

	var testRole1Name = "test role 1"

	role, err := r.RBAC.NewRole(tCtx, repo.NewRoleOpts{
		Name:      testRole1Name,
		UserType:  domain.TenantUserType,
		IsDefault: false,
	})

	require.NoError(t, err)
	assert.Equal(t, role.Name, testRole1Name)
	assert.Equal(t, role.TenantID, tnt.ID)

	rGet, err := r.RBAC.GetRole(tCtx, role.ID)

	require.NoError(t, err)
	assert.Equal(t, rGet.Name, testRole1Name)
	assert.Equal(t, rGet.TenantID, tnt.ID)
}

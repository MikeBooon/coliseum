package service_test

import (
	"testing"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/MikeBooon/coliseum/internal/tenant"
	"github.com/MikeBooon/coliseum/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRole(t *testing.T) {
	ctx := t.Context()
	s := service.NewServices(testDeps.System.DB)

	tnt, err := s.Tenant.New(ctx, "TestNewRole")

	require.NoError(t, err)

	tCtx := tenant.NewContext(ctx, tnt.ID)

	var testRole1Name = "test role 1"

	r, err := s.RBAC.NewRole(tCtx, testRole1Name, domain.ClientUserType)

	require.NoError(t, err)
	assert.Equal(t, r.Name, testRole1Name)
	assert.Equal(t, r.TenantID, tnt.ID)
}

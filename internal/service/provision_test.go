package service_test

import (
	"testing"

	"github.com/MikeBooon/coliseum/internal/service"
	"github.com/MikeBooon/coliseum/internal/tenant"
	"github.com/stretchr/testify/require"
)

func TestScaffoldTenant(t *testing.T) {
	ctx := t.Context()
	s := testDeps.System.Svcs

	tnt, err := s.Provision.ScaffoldTenant(ctx, service.ScaffoldTenantOpts{
		TenantName: "test",
		UserEmail:  "test@test.com",
	})

	require.NoError(t, err)

	_ = tenant.NewContext(ctx, tnt.ID)
}

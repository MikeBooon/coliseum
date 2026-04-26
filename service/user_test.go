package service_test

import (
	"testing"

	"github.com/MikeBooon/coliseum/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	ctx := t.Context()
	s := service.NewServices(testDeps.System)

	tenant, err := s.Tenant.New(ctx, "TestNewUser")

	require.NoError(t, err)

	u, err := s.User.GetByEmail(ctx, tenant.ID, "obama@test.com")

	require.ErrorIs(t, err, service.ErrUserNotFound)
	assert.Nil(t, u)

	u, err = s.User.New(ctx, tenant.ID, "obama@test.com")

	require.NoError(t, err)
	assert.NotNil(t, u)

	u, err = s.User.GetByEmail(ctx, tenant.ID, "obama@test.com")

	require.NoError(t, err)
	assert.NotNil(t, u)
}

package repo_test

// func TestNewUser(t *testing.T) {
// 	ctx := t.Context()
// 	s := repo.NewRepos(testDeps.System)

// 	tenant, err := s.Tenant.New(ctx, "TestNewUser")

// 	require.NoError(t, err)

// 	u, err := s.User.GetByEmail(ctx, tenant.ID, "obama@test.com")

// 	require.ErrorIs(t, err, repo.ErrUserNotFound)
// 	assert.Nil(t, u)

// 	u, err = s.User.New(ctx, tenant.ID, "obama@test.com")

// 	require.NoError(t, err)
// 	assert.NotNil(t, u)

// 	u, err = s.User.GetByEmail(ctx, tenant.ID, "obama@test.com")

// 	require.NoError(t, err)
// 	assert.NotNil(t, u)
// }

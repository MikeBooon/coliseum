package service

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/MikeBooon/coliseum/service/store"
)

type UserService struct {
	*store.Provider
}

func (s UserService) New(
	ctx context.Context,
	email string,
) (*dao.User, error) {
	tStore := s.NewTenantStore(ctx)
	u := &dao.User{Email: email}
	err := tStore.NewInsert(u).Scan(ctx)
	return u, err
}

func (s UserService) GetByEmail(
	ctx context.Context,
	email string,
) (*dao.User, error) {
	tStore := s.NewTenantStore(ctx)
	u := new(dao.User)
	err := tStore.NewSelect(u).
		Where("email = ?", email).Scan(ctx)

	if isNoRows(err) {
		return nil, ErrUserNotFound
	}

	return u, err
}

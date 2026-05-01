package service

import (
	"context"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/MikeBooon/coliseum/service/store"
	"github.com/google/uuid"
)

type UserService struct {
	*store.Provider
}

type NewUserOpts struct {
	Email  string
	Type   domain.UserType
	RoleID uuid.UUID
}

func (s UserService) New(
	ctx context.Context,
	opts NewUserOpts,
) (*dao.User, error) {
	tStore := s.NewTenantStore(ctx)
	u := &dao.User{Email: opts.Email, Type: opts.Type, RoleID: opts.RoleID}
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

func (s UserService) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*dao.User, error) {
	tStore := s.NewTenantStore(ctx)
	u := new(dao.User)
	err := tStore.NewSelect(u).
		Where("id = ?", id).Scan(ctx)

	if isNoRows(err) {
		return nil, ErrUserNotFound
	}

	return u, err
}

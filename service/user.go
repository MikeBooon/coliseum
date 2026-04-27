package service

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/db/dao"
	"github.com/google/uuid"
)

type UserService struct {
	db db.IDB
}

func (s UserService) New(ctx context.Context, tenantID uuid.UUID, email string) (*dao.User, error) {
	u := &dao.User{Email: email}
	err := s.db.NewInsert().Model(u).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s UserService) GetByEmail(ctx context.Context, tenantID uuid.UUID, email string) (*dao.User, error) {
	u := new(dao.User)
	err := s.db.NewSelect().Model(u).Apply(whereTenant(tenantID)).
		Where("email = ?", email).Scan(ctx)

	if isNoRows(err) {
		return nil, ErrUserNotFound
	}

	return u, err
}

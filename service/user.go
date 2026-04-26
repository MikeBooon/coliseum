package service

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/db/dao"
)

type UserService struct {
	db *db.DB
}

func (s UserService) New(ctx context.Context, email string) (*dao.User, error) {
	user := &dao.User{Email: email}
	_, err := s.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

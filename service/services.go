package service

import "github.com/MikeBooon/coliseum/internal/system"

type Services struct {
	User UserService
}

func NewServices(sys system.System) *Services {
	s := Services{User: UserService{db: sys.DB}}
	return &s
}

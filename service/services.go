package service

import "github.com/MikeBooon/coliseum/internal/system"

type Services struct {
	User   UserService
	Tenant TenantService
}

func NewServices(sys system.System) *Services {
	return &Services{
		User:   UserService{db: sys.DB},
		Tenant: TenantService{db: sys.DB},
	}
}

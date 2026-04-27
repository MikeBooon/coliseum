package service

import (
	"github.com/MikeBooon/coliseum/internal/db"
)

type Services struct {
	User   UserService
	Tenant TenantService
	RBAC   RBACService
}

func NewServices(db db.IDB) *Services {
	return &Services{
		User:   UserService{db},
		Tenant: TenantService{db},
		RBAC:   RBACService{db},
	}
}

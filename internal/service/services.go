package service

import (
	"github.com/MikeBooon/coliseum/internal/repo"
)

type Services struct {
	Provision *ProvisionService
}

func NewServices(repos *repo.Repos) *Services {
	return &Services{
		Provision: &ProvisionService{repos: repos},
	}
}

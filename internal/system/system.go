package system

import (
	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/internal/repo"
	"github.com/MikeBooon/coliseum/internal/service"
)

type System struct {
	DB        *db.DB
	EnvConfig *config.EnvConfig
	Svcs      *service.Services
	Repos     *repo.Repos
}

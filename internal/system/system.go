package system

import (
	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/MikeBooon/coliseum/internal/db"
	"github.com/MikeBooon/coliseum/service"
)

type System struct {
	DB     *db.DB
	Config *config.Config
	Svcs   *service.Services
}

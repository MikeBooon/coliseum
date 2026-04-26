package system

import (
	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/MikeBooon/coliseum/internal/db"
)

type System struct {
	DB     *db.DB
	Config *config.Config
}

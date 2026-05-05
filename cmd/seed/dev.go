package main

import (
	"github.com/MikeBooon/coliseum/internal/system"
	"github.com/MikeBooon/coliseum/service"
)

func devSeed(sys system.System) error {
	_ = service.NewServices(sys.DB)

	return nil
}

func init() {
	seedFiles["dev"] = devSeed
}

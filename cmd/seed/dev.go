package main

import (
	"errors"

	"github.com/MikeBooon/coliseum/internal/system"
	"github.com/MikeBooon/coliseum/service"
)

func devSeed(sys system.System) error {
	_ = service.NewServices(sys.DB)

	return errors.New("test error")
}

func init() {
	seedFiles["dev"] = devSeed
}

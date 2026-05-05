package main

import (
	"github.com/MikeBooon/coliseum/internal/system"
)

func devSeed(sys system.System) error {

	return nil
}

func init() {
	seedFiles["dev"] = devSeed
}

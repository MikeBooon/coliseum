package main

import (
	"log/slog"

	"github.com/MikeBooon/coliseum/internal/rest"
)

func main() {
	slog.Info("Starting rest server")

	r := rest.NewRest()
	err := r.Start()
	if err != nil {
		panic(err)
	}
}

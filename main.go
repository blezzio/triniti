package main

import (
	"log/slog"
	"os"

	"github.com/blezzio/triniti/data/repositories"
	"github.com/blezzio/triniti/infra"
	"github.com/blezzio/triniti/services/usecases"
)

func main() {
	build()
}

func build() {
	conn, err := infra.NewPostgresConn()
	if err != nil {
		slog.Error("failed to create db with error: %v", err)
		os.Exit(1)
	}
	repo := repositories.NewURL(conn)
	_ = usecases.NewURL(repo, nil)
}

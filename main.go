package main

import (
	"crypto/sha512"
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
	hash := usecases.NewHasher(sha512.New())
	_ = usecases.NewURL(repo, hash)
}

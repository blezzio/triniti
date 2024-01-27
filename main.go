package main

import (
	"context"
	"crypto/sha512"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blezzio/triniti/data/repositories"
	"github.com/blezzio/triniti/handlers/routers"
	"github.com/blezzio/triniti/infra"
	"github.com/blezzio/triniti/middlewares"
	"github.com/blezzio/triniti/services/usecases"
)

func main() {
	server, teardown := build()

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		slog.Error("server encounter error, tearing down...", "error", err)
		teardown()
	case sig := <-quit:
		slog.Warn("system signal captured, tearing down...", "sginal", sig)
		teardown()
	}
	slog.Info("teared down")
}

func build() (server *infra.Server, teardown func()) {
	conn, err := infra.NewPostgresConn()
	if err != nil {
		slog.Error("failed to create db", "error", err)
		os.Exit(1)
	}
	repo := repositories.NewURL(conn)
	hash := usecases.NewHasher(sha512.New())
	uc := usecases.NewURL(repo, hash)
	router := routers.NewURL(uc)

	reql := middlewares.NewReqLogger()

	server = infra.NewServer(
		infra.WithRouter(router),
		infra.WithMiddleware(reql),
	)
	teardown = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := conn.Close(); err != nil {
			slog.WarnContext(ctx, "failed to close connection", "error", err)
		}
		slog.Info("db connection closed")

		if err := server.Shutdown(ctx); err != nil {
			slog.WarnContext(ctx, "graceful shutdown server failed", "error", err)
		}
		slog.Info("server shutdown")
	}

	return
}

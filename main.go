package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blezzio/triniti/apis/handlers"
	"github.com/blezzio/triniti/assets"
	"github.com/blezzio/triniti/data/repositories"
	"github.com/blezzio/triniti/infra"
	"github.com/blezzio/triniti/middlewares"
	_ "github.com/blezzio/triniti/presentation/l10n"
	"github.com/blezzio/triniti/presentation/views"
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
	hasher := usecases.NewHasher()
	uc := usecases.NewURL(repo, hasher)

	indexView := views.NewIndex(assets.FS)
	successView := views.NewSucccess(assets.FS)

	handler := handlers.NewURL(
		uc,
		handlers.WithView(indexView, handlers.IndexView),
		handlers.WithView(successView, handlers.SuccessView),
		handlers.WithView(indexView, handlers.ErrorView),
	)

	reql := middlewares.NewReqLogger()
	respcom := middlewares.NewRespCompressor()
	static := middlewares.NewStatic(assets.FS, "static")
	favico := middlewares.NewFavIco(assets.FS, "static")

	server = infra.NewServer(
		infra.WithHandler(handler),
		infra.WithMiddleware(static),
		infra.WithMiddleware(reql),
		infra.WithMiddleware(respcom),
		infra.WithMiddleware(favico),
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

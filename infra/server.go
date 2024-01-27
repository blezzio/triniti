package infra

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/blezzio/triniti/utils"
)

type Server struct {
	srv      *http.Server
	serveMux *http.ServeMux
	logger   Logger
}

func NewServer(opts ...ServerOpt) *Server {
	h := http.NewServeMux()
	s := &Server{
		srv: &http.Server{
			Addr:    os.Getenv("PORT"),
			Handler: h,
		},
		serveMux: h,
		logger:   slog.Default(),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) ListenAndServe() error {
	s.logger.Info("server starts serving at", "address", s.srv.Addr)
	return utils.Trace(s.srv.ListenAndServe(), "server abort")
}

func (s *Server) Shutdown(ctx context.Context) error {
	return utils.Trace(s.srv.Shutdown(ctx), "shutdown server failed")
}

package infra

import (
	"net/http"
)

type handler interface {
	Build(*http.ServeMux)
}

type ServerOpt func(*Server)

func WithHandler(r handler) ServerOpt {
	return func(s *Server) {
		r.Build(s.serveMux)
	}
}

func WithAddr(addr string) ServerOpt {
	return func(s *Server) {
		s.srv.Addr = addr
	}
}

func WithLogger(logger Logger) ServerOpt {
	return func(server *Server) {
		server.logger = logger
	}
}

func WithMiddleware(mw Middleware) ServerOpt {
	return func(s *Server) {
		s.srv.Handler = mw.Wrap(s.srv.Handler)
	}
}

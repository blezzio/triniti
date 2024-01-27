package infra

import (
	"net/http"
)

type router interface {
	Route(*http.ServeMux)
}

type ServerOpt func(*Server)

func WithRouter(r router) ServerOpt {
	return func(s *Server) {
		r.Route(s.serveMux)
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

package infra

import "net/http"

type ServerOpt func(*http.Server)

func WithRouter(r router) ServerOpt {
	return func(s *http.Server) {
		r.Route(s.Handler.(*http.ServeMux))
	}
}

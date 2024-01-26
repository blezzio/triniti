package infra

import (
	"net/http"
	"os"
)

type router interface {
	Route(*http.ServeMux)
}

func NewServer(opts ...ServerOpt) *http.Server {
	s := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: http.NewServeMux(),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

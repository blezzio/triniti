package middlewares

import (
	"net/http"
	"time"
)

type reqlmsg struct {
	method  string
	uri     string
	latency time.Duration
	status  int
}

type reqlrec struct {
	http.ResponseWriter
	status    int
	startedAt time.Time
}

func (r *reqlrec) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

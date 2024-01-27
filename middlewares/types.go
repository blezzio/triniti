package middlewares

import (
	"io"
	"net/http"
	"time"
)

type writerConstructor func(w io.Writer) (writer io.WriteCloser, encoding string, err error)

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

type compressWriter struct {
	http.ResponseWriter
	writer io.WriteCloser
}

func (r *compressWriter) Write(b []byte) (int, error) {
	if length, err := r.writer.Write(b); err == nil {
		return length, nil
	}
	return r.ResponseWriter.Write(b)
}

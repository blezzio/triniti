package middlewares

import (
	"compress/gzip"
	"log/slog"
	"net/http"
)

type RespCompressor struct {
	wconstr writerConstructor
	handler http.Handler
	logger  Logger
}

func NewRespCompressor(opts ...RespCompressorOpt) *RespCompressor {
	mw := &RespCompressor{logger: slog.Default()}
	WithGZip(gzip.BestCompression)(mw)

	for _, opt := range opts {
		opt(mw)
	}
	return mw
}

func (mw *RespCompressor) Wrap(h http.Handler) http.Handler {
	return &RespCompressor{
		handler: h,
		wconstr: mw.wconstr,
	}
}

func (mw *RespCompressor) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	compw, contentEncoding, err := mw.wconstr(rw)
	if err != nil {
		mw.logger.Warn("failed to create compression writer", "error", err)
	} else {
		defer compw.Close()
		rw.Header().Add("Content-Encoding", contentEncoding)
	}
	mw.handler.ServeHTTP(
		&compressWriter{
			ResponseWriter: rw,
			writer:         compw,
		},
		req,
	)
}

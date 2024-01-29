package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type ReqLogger struct {
	handler http.Handler
	logger  Logger
}

func NewReqLogger(opts ...ReqLoggerOpt) *ReqLogger {
	mw := &ReqLogger{
		logger: slog.Default(),
	}
	for _, opt := range opts {
		opt(mw)
	}
	return mw
}

func (mw *ReqLogger) Wrap(h http.Handler) http.Handler {
	return &ReqLogger{
		handler: h,
		logger:  mw.logger,
	}
}

func (mw *ReqLogger) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	recorder := &reqlrec{
		ResponseWriter: rw,
		startedAt:      time.Now(),
		status:         http.StatusOK,
	}
	mw.handler.ServeHTTP(recorder, req)
	mw.log(mw.buildMsg(req, recorder))
}

func (mw *ReqLogger) log(msg *reqlmsg) {
	latencyUnit := "s"
	latency := int64(msg.latency.Seconds())
	if latency == 0 {
		latencyUnit = "ms"
		latency = msg.latency.Milliseconds()
	}
	if latency == 0 {
		latencyUnit = "μs"
		latency = msg.latency.Microseconds()
	}

	if msg.status > 399 {
		mw.logger.Error(
			"request failed",
			"uri", msg.uri,
			"method", msg.method,
			"status", msg.status,
			"latency", fmt.Sprintf("%dms", msg.latency.Milliseconds()),
		)
	} else {
		mw.logger.Info(
			"request succeeded",
			"uri", msg.uri,
			"method", msg.method,
			"status", msg.status,
			"latency", fmt.Sprintf("%d%s", latency, latencyUnit),
		)
	}
}

func (mw *ReqLogger) buildMsg(req *http.Request, rec *reqlrec) *reqlmsg {
	return &reqlmsg{
		method:  req.Method,
		status:  rec.status,
		latency: time.Since(rec.startedAt),
		uri:     req.RequestURI,
	}
}

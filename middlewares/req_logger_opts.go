package middlewares

type ReqLoggerOpt func(*ReqLogger)

func WithLogger(logger Logger) ReqLoggerOpt {
	return func(reql *ReqLogger) {
		reql.logger = logger
	}
}

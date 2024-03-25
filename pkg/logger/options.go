package logger

func (l *logger) InfoRequest(method, path, time string) {
	l.log.Log().
		Str("level", "new request").
		Str("method", method).
		Str("request time", time).
		Send()
}

func (l *logger) Error(err any) {
	l.log.Error().Any("error", err).Send()
}

func (l *logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

package logger

import "fmt"

func (l *logger) Info(msg string, args ...any) {
	l.log.Info().Msg(fmt.Sprintf(msg, args...))
}

func (l *logger) Error(msg string, args ...any) {
	l.log.Error().Msg(fmt.Sprintf(msg, args...))
}

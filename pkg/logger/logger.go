package logger

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

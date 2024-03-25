package logger

type Logger interface {
	Info(msg string)
	InfoRequest(method, path, time string)
	Error(err any)
}

package logger

import "github.com/rs/zerolog"

type logger struct {
	log zerolog.Logger
}

func New(log zerolog.Logger) Logger {
	return &logger{log}
}

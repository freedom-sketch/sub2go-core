package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	*log.Logger
	file *os.File
}

func (l *Logger) Close() error {
	return l.file.Close()
}

func NewLogger(path string, level string) (*Logger, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	l := log.New()
	l.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     false,
	})

	lvl, err := log.ParseLevel(level)
	if err != nil {
		lvl = log.InfoLevel
	}
	l.SetLevel(lvl)
	l.SetOutput(io.MultiWriter(os.Stdout, f))

	return &Logger{
		Logger: l,
		file:   f,
	}, nil
}

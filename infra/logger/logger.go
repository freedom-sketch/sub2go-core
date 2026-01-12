package logger

import (
	"io"
	"os"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	*log.Logger
	file *os.File
}

func (l *Logger) Close() error {
	return l.file.Close()
}

func NewLogger(cfg *config.Logging) (*Logger, error) {
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(cfg.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	l := log.New()

	l.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     false,
	})

	lvl, err := log.ParseLevel(cfg.Level)
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

package logger

import (
	"io"
	"os"

	"github.com/freedom-sketch/sub2go-core/internal/config"
	log "github.com/sirupsen/logrus"
)

var Log *log.Logger
var logFile *os.File

func Init(cfg *config.Logging) error {
	if err := os.MkdirAll("logs", 0755); err != nil {
		return err
	}

	f, err := os.OpenFile("logs/"+cfg.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	logFile = f

	Log = log.New()

	Log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     false,
	})

	lvl, err := log.ParseLevel(cfg.Level)
	if err != nil {
		lvl = log.InfoLevel
	}
	Log.SetLevel(lvl)
	Log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	return nil
}

func Close() error {
	if logFile == nil {
		return nil
	}
	err := logFile.Close()
	logFile = nil

	return err
}

package utils

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"
)

func GetFileToSave(logName string) *os.File {
	logsDir := os.Getenv("LOGS_DIR")
	now := time.Now()
	date := now.Format("2006_01_02")
	name := fmt.Sprintf("%s/%s_log_%s.log", logsDir, logName, date)

	CreateIfNotExists(name)
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error on open log file: %s", err)
	}

	return f
}

func CreateFolderIfNotExists() {
	logsDir := os.Getenv("LOGS_DIR")
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		os.Mkdir(logsDir, 0755)
	}
}

func CreateIfNotExists(name string) error {
	logExists := true

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			logExists = false
		}
	}
	if logExists {
		return nil
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer func() {
		f.Close()
	}()

	return nil
}

func GetLogLevel(status int) slog.Level {
	switch {
	case status >= 500:
		return slog.LevelError
	case status >= 400:
		return slog.LevelWarn
	default:
		return slog.LevelInfo
	}
}

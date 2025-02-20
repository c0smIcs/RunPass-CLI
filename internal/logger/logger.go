package logger

import (
	"log/slog"
	"os"
)

// SetupLogger настраивает и возвращает JSON-логгер
func SetupLogger(logFile string) *slog.Logger {
	var output *os.File
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		output = file
	} else {
		output = os.Stdout
	}

	logger := slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	return logger
}
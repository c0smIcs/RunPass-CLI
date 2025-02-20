package main

import (
	"log/slog"
	"os"

	"github.com/c0smIcs/RanPass_CLI/cmd"
	"github.com/c0smIcs/RanPass_CLI/internal/logger"
)

func main() {
	logger := logger.SetupLogger("app.log")
	slog.SetDefault(logger)

	if err := cmd.RootCmd.Execute(); err != nil {
		slog.Error("Ошибка выполнения команды", "error", err)
		os.Exit(1)
	}
}

/*
	Данный файл предназначен для входа в приложение.

	Что происходит в этом файле:
	1. Создает логгер, который пишет в "app.log" (в формате JSON)
	2. Устанавливает его как логгер по умолчанию для приложения
	3. 
*/

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

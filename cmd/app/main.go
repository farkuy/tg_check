package main

import (
	"log/slog"
	"os"
	"tg_check/internal/config"
	"tg_check/internal/logger"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		slog.Error("Ошибка инициализации конфига", "error", err)
		os.Exit(1)
	}

	logger.Init(cfg.Environment)

	//подрубиться к бд

	//написать основные ручки

	//запуск сервера
}

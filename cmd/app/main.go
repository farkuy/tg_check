package main

import (
	"fmt"
	"log/slog"
	"os"
	"tg_check/internal/config"
	"tg_check/internal/database"
	"tg_check/internal/logger"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		slog.Error("Ошибка инициализации конфига", err)
		os.Exit(1)
	}

	logger.Init(cfg.Environment)

	bd, err := database.Init(cfg.DatabasePath)
	if err != nil {
		slog.Error("Ошибка подключение к бд", err)
		os.Exit(1)
	}

	fmt.Println(bd)
	//написать основные ручки

	//запуск сервера
}

package main

import (
	"fmt"
	"log/slog"
	"os"
	"tg_check/internal/config"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	fmt.Println(cfg)
	//сделать логер

	//подрубиться к бд

	//написать основные ручки

	//запуск сервера
}

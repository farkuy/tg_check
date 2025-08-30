package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"tg_check/internal/config"
	"tg_check/internal/database"
	targets "tg_check/internal/domain/storages"
	"tg_check/internal/logger"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		slog.Error("Ошибка инициализации конфига", err)
		os.Exit(1)
	}

	logger.Init(cfg.Environment)

	storage, err := database.Init(cfg.DatabasePath)
	if err != nil {
		slog.Error("Ошибка подключение к бд", err)
		os.Exit(1)
	}

	fmt.Println(storage.DB)
	router := chi.NewRouter()

	router.Post("/storages", targets.PostTarget(storage))

	http.ListenAndServe()
	//написать основные ручки

	//запуск сервера
}

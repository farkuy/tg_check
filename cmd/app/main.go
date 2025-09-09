package main

import (
	"log/slog"
	"net/http"
	"os"
	"tg_check/internal/config"
	"tg_check/internal/database"
	storageHistoty "tg_check/internal/domain/storageHistory"
	"tg_check/internal/domain/storages"
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

	router := chi.NewRouter()

	storages.StoragesHandlersInit(router, storage)
	storageHistoty.StoragesHistoryHandlersInit(router, storage)

	server := &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTime,
		WriteTimeout: cfg.SendTime,
	}

	if err = server.ListenAndServe(); err != nil {
		slog.Error("Ошибка старта сервера:", err)
	}
	slog.Info("Сервер стартовал на:", cfg.Host, cfg.Port)
}

package storages

import (
	"tg_check/internal/database"

	"github.com/go-chi/chi/v5"
)

func StoragesHandlersInit(router *chi.Mux, storage *database.Storage) {
	wrapper := &StorageWrapper{storage}

	router.Post("/storages", postStorage(wrapper))
}

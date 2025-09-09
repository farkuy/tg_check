package storages

import (
	"tg_check/internal/database"

	"github.com/go-chi/chi/v5"
)

func StoragesHandlersInit(router *chi.Mux, storage *database.Storage) {
	wrapperStorage := &StorageWrapper{storage}

	router.Post("/storage", postStorage(wrapperStorage))
	router.Get("/storage", getStorage(wrapperStorage))
	router.Put("/storage", updateStorage(wrapperStorage))
	router.Delete("/storage", delStorage(wrapperStorage))
}

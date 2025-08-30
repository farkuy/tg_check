package targets

import (
	"tg_check/internal/database"

	"github.com/go-chi/chi"
)

func TargetsHandlers(router *chi.Mux, storage *database.Storage) {
	router.Post("/storages", PostTarget(postTarget(storage)))

}

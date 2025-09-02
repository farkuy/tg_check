package storages

import (
	"log/slog"
	"net/http"
	"tg_check/internal/httpModel"
	"time"

	"github.com/go-chi/render"
)

type requestStorage struct {
	Sum          int       `json:"sum" validate:"required"`
	Accumulated  int       `json:"accumulated" validate:"required"`
	DeadLineDate time.Time `json:"deadLineDate" validate:"required"`
}

type responseStorage struct {
	Id int `json:"id"`
	httpModel.Response
}

type storageCreacte interface {
	postTargetSql(sum, accumulated int, deadLineDate time.Time) (*Storage, error)
}

func postStorage(tCreate storageCreacte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := slog.With("service", "storage")

		var req requestStorage
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Ошибка получения данных из json", err)
			render.JSON(w, r, httpModel.Error("не правильно переданны данные"))
			return
		}

		err = validateRequestPost(req, log, w, r)
		if err != nil {
			render.JSON(w, r, httpModel.Error("не правильно переданны данные"))
			return
		}

		res, err := tCreate.postTargetSql(req.Sum, req.Accumulated, req.DeadLineDate)
		if err != nil {
			log.Error("Ошибка создания storage", err)
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при создании цели"))
			return
		}

		if res == nil {
			log.Error("Ошибка получение данных storage: пустой рузельтат ответа")
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при создании цели"))
			return
		}

		render.JSON(w, r, responseStorage{
			Id:       res.Id,
			Response: httpModel.OK(),
		})
	}
}

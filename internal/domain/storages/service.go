package storages

import (
	"log/slog"
	"net/http"
	"strconv"
	"tg_check/internal/httpModel"
	"time"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type requestStorageCreate struct {
	Sum          int       `json:"sum" validate:"required"`
	Accumulated  int       `json:"accumulated" validate:"required"`
	DeadLineDate time.Time `json:"deadLineDate" validate:"required"`
}

type requestStorageFind struct {
	Id int `json:"id" validate:"required"`
}

type responseStorage struct {
	Storage *Storage `json:"storage"`
	httpModel.Response
}

type storageCreacte interface {
	postStorageSql(sum, accumulated int, deadLineDate time.Time) (*Storage, error)
}

type storageFind interface {
	getStorageSql(id int) (*Storage, error)
}

func postStorage(sCreate storageCreacte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := slog.With("service", "storage post")

		var req requestStorageCreate
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Ошибка получения данных из json", err)
			render.JSON(w, r, httpModel.Error("не правильно переданны данные"))
			return
		}

		err = validateRequestPost(req, log, w, r)
		if err != nil {
			log.Error("Ошибка валидации данных из json", err)
			render.JSON(w, r, httpModel.Error("ошибка валидации данных"))
			return
		}

		res, err := sCreate.postStorageSql(req.Sum, req.Accumulated, req.DeadLineDate)
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
			Storage:  res,
			Response: httpModel.OK(),
		})
	}
}

func getStorage(sFind storageFind) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := slog.With("service", "storage get")

		var req requestStorageFind
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Ошибка получения данных из json", err)
			render.JSON(w, r, httpModel.Error("не правильно переданны данные"))
			return
		}

		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			errors := err.(validator.ValidationErrors)
			log.Info("Переданны не все данные", errors)
			render.JSON(w, r, httpModel.Error("не правильно переданны данные"))
			return
		}

		query := r.URL.Query()
		id := query.Get("id")
		if id == "" {
			log.Error("Не найден параметр id для запроса. query: ", query)
			render.JSON(w, r, httpModel.Error("Не задан id"))
			return
		}

		idNum, err := strconv.Atoi(id)
		if err != nil {
			log.Error("Ошибка преобразования id к числу ", query)
			render.JSON(w, r, httpModel.Error("Ошибка с заданным id"))
			return
		}

		res, err := sFind.getStorageSql(idNum)
		if err != nil {
			log.Error("Ошибка нахождения storage", err)
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при нахождении цели"))
			return
		}

		if res == nil {
			log.Error("Ошибка получение данных storage: пустой рузельтат ответа")
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при нахождении цели"))
			return
		}

		render.JSON(w, r, responseStorage{
			Storage:  res,
			Response: httpModel.OK(),
		})
	}
}

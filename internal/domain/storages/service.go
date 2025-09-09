package storages

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	storageHistoty "tg_check/internal/domain/storageHistory"
	"tg_check/internal/httpModel"
	"time"

	"github.com/go-chi/render"
)

type requestStorage struct {
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

type storageCreate interface {
	postStorageSql(sum, accumulated int, deadLineDate time.Time) (*Storage, error)
}

type storageFind interface {
	getStorageSql(id int) (*Storage, error)
}

type storageUpdate interface {
	updateStorageSql(id, sum, accumulated int, deadLineDate time.Time) (*Storage, error)
}

type storageDelete interface {
	deleteStorageSql(id int) error
}

func postStorage(sCreate storageCreate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := slog.With("service", "storage post")

		var req requestStorage
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Ошибка получения данных из json", err)
			render.JSON(w, r, httpModel.Error("не правильно переданны данные"))
			return
		}

		err = validateRequestPost(req, log, w, r)
		if err != nil {
			log.Error("Ошибка валидации данных из json", err)
			render.JSON(w, r, httpModel.Error(err.Error()))
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

		storageHistoty.StoragePoint.Update(storageHistoty.CreateHistory{TypeChange: "pluse", ChangeSum: res.Accumulated, StorageId: res.Id})
	}
}

func getStorage(sFind storageFind) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := slog.With("service", "storage get")

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
			if errors.Is(sql.ErrNoRows, err) {
				render.JSON(w, r, httpModel.Error("Цели с такой id не существует"))
				return
			}
			log.Error("Ошибка нахождения storage", err)
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при нахождении цели"))
			return
		}

		if res == nil {
			log.Error("Ошибка получение данных storage: пустой рузельтат ответа")
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при нахождении цели"))
			return
		}

		fmt.Println(res)

		render.JSON(w, r, responseStorage{
			Storage:  res,
			Response: httpModel.OK(),
		})
	}
}

func updateStorage(sUpdate storageUpdate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := slog.With("service", "storage put")

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

		var req requestStorage
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Ошибка получения данных из json", err)
			render.JSON(w, r, httpModel.Error("не правильно переданны данные"))
			return
		}

		err = validateRequestPost(req, log, w, r)
		if err != nil {
			log.Error("Ошибка валидации данных из json", err)
			render.JSON(w, r, httpModel.Error(err.Error()))
			return
		}

		res, err := sUpdate.updateStorageSql(idNum, req.Sum, req.Accumulated, req.DeadLineDate)
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				render.JSON(w, r, httpModel.Error("Цели с такой id не существует"))
				return
			}
			log.Error("Ошибка обновдения storage", err)
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при обновлении цели"))
			return
		}

		if res == nil {
			log.Error("Ошибка получения обновленных данных storage: пустой рузельтат ответа")
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при обновлении цели"))
			return
		}

		render.JSON(w, r, responseStorage{
			Storage:  res,
			Response: httpModel.OK(),
		})

		storageHistoty.StoragePoint.Update(storageHistoty.CreateHistory{TypeChange: "pluse", ChangeSum: res.Accumulated, StorageId: res.Id})
	}
}

func delStorage(sDelete storageDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := slog.With("service", "storage delete")

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

		err = sDelete.deleteStorageSql(idNum)
		if err != nil {
			log.Error("Ошибка удаления storage", err)
			render.JSON(w, r, httpModel.Error("Что-то пошло не так при удалении цели"))
			return
		}

		render.JSON(w, r, httpModel.OK())
	}
}

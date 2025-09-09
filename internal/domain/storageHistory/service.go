package storageHistoty

import (
	"log/slog"
)

type CreateHistory struct {
	TypeChange string
	ChangeSum  int
	StorageId  int
}

type storageHistotyPointCreate interface {
	postStorageHistotyPointSql(typeChange string, changeSum, storageId int) (*StorageHistotyPoint, error)
}

func PostStorageHistoryPoint(hCreate storageHistotyPointCreate, data CreateHistory) {
	log := slog.With("service", "storage history pint post")

	res, err := hCreate.postStorageHistotyPointSql(data.TypeChange, data.ChangeSum, data.StorageId)
	if err != nil {
		log.Error("Ошибка создания точки истории", err)
		return
	}

	if res == nil {
		log.Error("Ошибка получение данных точки истории: пустой рузельтат ответа")
		return
	}
}

// func getStorage(sFind storageFind) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log := slog.With("service", "storage get")

// 		query := r.URL.Query()
// 		id := query.Get("id")
// 		if id == "" {
// 			log.Error("Не найден параметр id для запроса. query: ", query)
// 			render.JSON(w, r, httpModel.Error("Не задан id"))
// 			return
// 		}

// 		idNum, err := strconv.Atoi(id)
// 		if err != nil {
// 			log.Error("Ошибка преобразования id к числу ", query)
// 			render.JSON(w, r, httpModel.Error("Ошибка с заданным id"))
// 			return
// 		}

// 		res, err := sFind.getStorageSql(idNum)
// 		if err != nil {
// 			if errors.Is(sql.ErrNoRows, err) {
// 				render.JSON(w, r, httpModel.Error("Цели с такой id не существует"))
// 				return
// 			}
// 			log.Error("Ошибка нахождения storage", err)
// 			render.JSON(w, r, httpModel.Error("Что-то пошло не так при нахождении цели"))
// 			return
// 		}

// 		if res == nil {
// 			log.Error("Ошибка получение данных storage: пустой рузельтат ответа")
// 			render.JSON(w, r, httpModel.Error("Что-то пошло не так при нахождении цели"))
// 			return
// 		}

// 		fmt.Println(res)

// 		render.JSON(w, r, responseStorage{
// 			Storage:  res,
// 			Response: httpModel.OK(),
// 		})
// 	}
// }

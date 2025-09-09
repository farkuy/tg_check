package storageHistoty

import (
	"tg_check/internal/database"

	"github.com/go-chi/chi/v5"
)

type StorageHistotyPointObserv struct {
	data    *CreateHistory
	hCreate *StorageHistoryWrapper
}

func (pointHistory *StorageHistotyPointObserv) Update(newData CreateHistory) {
	pointHistory.data = &newData

	defer pointHistory.PostStorage()
}

func (pointHistory *StorageHistotyPointObserv) PostStorage() {
	pointHistory.hCreate.postStorageHistotyPointSql(pointHistory.data.TypeChange, pointHistory.data.ChangeSum, pointHistory.data.StorageId)
}

func (pointHistory *StorageHistotyPointObserv) Clear() {
	pointHistory.data = nil
}

var StoragePoint = StorageHistotyPointObserv{data: nil, hCreate: nil}

func StoragesHistoryHandlersInit(router *chi.Mux, storage *database.Storage) {
	wrapperStorage := &StorageHistoryWrapper{storage}
	StoragePoint.hCreate = wrapperStorage
}

package storageHistoty

import (
	"log/slog"
	"tg_check/internal/database"
	"time"
)

type StorageHistotyPoint struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	ChangeSum int       `json:"changeSum"`
	Date      time.Time `json:"date"`
	StorageId int       `json:"storageId"`
}

type StorageHistoryWrapper struct {
	*database.Storage
}

func (storage *StorageHistoryWrapper) postStorageHistotyPointSql(typeChange string, changeSum, storageId int) (*StorageHistotyPoint, error) {
	log := slog.With("repository", "storage hisory point post")

	data := &StorageHistotyPoint{}
	err := storage.DB.QueryRow(postStorageHistoryQuery, typeChange, changeSum, storageId).Scan(
		&data.Id,
		&data.Type,
		&data.ChangeSum,
		&data.Date,
		&data.StorageId,
	)

	if err != nil {
		log.Error("ошибка создания точки истории", err)
		return nil, err
	}

	return data, nil
}

// func (storage *StorageHistoryWrapper) getStorageSql(id int) (*Storage, error) {
// 	log := slog.With("repository", "storage get")

// 	data := &Storage{}
// 	err := storage.DB.QueryRow(getStorageQuery, id).Scan(
// 		&data.Id,
// 		&data.Sum,
// 		&data.Accumulated,
// 		&data.CreateDate,
// 		&data.DeadLineDate,
// 	)

// 	if err != nil {
// 		log.Error("ошибка получение цели", err)
// 		return nil, err
// 	}

// 	return data, nil
// }

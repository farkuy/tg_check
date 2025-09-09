package storages

import (
	"log/slog"
	"tg_check/internal/database"
	"time"
)

type Storage struct {
	Id           int       `json:"id"`
	Sum          int       `json:"sum"`
	Accumulated  int       `json:"accumulated"`
	CreateDate   time.Time `json:"createDate"`
	DeadLineDate time.Time `json:"deadLineDate"`
}

type StorageWrapper struct {
	*database.Storage
}

func (storage *StorageWrapper) postStorageSql(sum, accumulated int, deadLineDate time.Time) (*Storage, error) {
	log := slog.With("repository", "storage post")

	data := &Storage{}
	err := storage.DB.QueryRow(postStorageQuery, sum, accumulated, deadLineDate).Scan(
		&data.Id,
		&data.Sum,
		&data.Accumulated,
		&data.CreateDate,
		&data.DeadLineDate,
	)

	if err != nil {
		log.Error("ошибка создания", err)
		return nil, err
	}

	return data, nil
}

func (storage *StorageWrapper) getStorageSql(id int) (*Storage, error) {
	log := slog.With("repository", "storage get")

	data := &Storage{}
	err := storage.DB.QueryRow(getStorageQuery, id).Scan(
		&data.Id,
		&data.Sum,
		&data.Accumulated,
		&data.CreateDate,
		&data.DeadLineDate,
	)

	if err != nil {
		log.Error("ошибка получение цели", err)
		return nil, err
	}

	return data, nil
}

func (storage *StorageWrapper) updateStorageSql(id, sum, accumulated int, deadLineDate time.Time) (*Storage, error) {
	log := slog.With("repository", "storage put")

	data := &Storage{}
	err := storage.DB.QueryRow(updateStorageSum, sum, accumulated, deadLineDate, id).Scan(
		&data.Id,
		&data.Sum,
		&data.Accumulated,
		&data.CreateDate,
		&data.DeadLineDate,
	)

	if err != nil {
		log.Error("ошибка обновления", err)
		return nil, err
	}

	return data, nil
}

func (storage *StorageWrapper) deleteStorageSql(id int) error {
	log := slog.With("repository", "storage delete")

	_, err := storage.DB.Exec(deleteStorage, id)

	if err != nil {
		log.Error("ошибка получение цели", err)
		return err
	}

	return nil
}

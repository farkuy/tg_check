package storages

import (
	"fmt"
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

func (storage *StorageWrapper) postTargetSql(sum, accumulated int, deadLineDate time.Time) (*Storage, error) {
	log := slog.With("repository", "storage")

	target := &Storage{}
	err := storage.DB.QueryRow(postStorageQuery, sum, accumulated, deadLineDate).Scan(
		&target.Id,
		&target.Sum,
		&target.Accumulated,
		&target.CreateDate,
		&target.DeadLineDate,
	)

	if err != nil {
		log.Error("ошибка создания", err)
		return nil, fmt.Errorf("ошибка создания target с sum:(%v), accumulated:(%v), deadLineDate:(%v)", sum, accumulated, deadLineDate)
	}

	return target, nil
}

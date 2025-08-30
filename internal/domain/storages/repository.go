package targets

import (
	"fmt"
	"log/slog"
	"tg_check/internal/database"
	"time"
)

type Target struct {
	Id           int       `json:"id"`
	Sum          int       `json:"sum"`
	Accumulated  int       `json:"accumulated"`
	DeadLineDate time.Time `json:"deadLineDate"`
}

func postTarget(storage *database.Storage, sum, accumulated int, deadLineDate time.Time) (*Target, error) {
	log := slog.With("repository", "target", "create")

	target := &Target{}
	err := storage.DB.QueryRow(postTargetQuery, sum, accumulated, deadLineDate).Scan(
		&target.Id,
		&target.Sum,
		&target.Accumulated,
		&target.DeadLineDate,
	)
	if err != nil {
		log.Error(err.Error())
		return nil, fmt.Errorf("ошибка создания target с sum:(%v), accumulated:(%v), deadLineDate:(%v)", sum, accumulated, deadLineDate)
	}

	return target, nil
}

func getTarget(storage *database.Storage, id int) (Target, error) {
	log := slog.With("repository", "target", "get")

	var target Target
	err := storage.DB.QueryRow(getTargetQuery, id).Scan(&target.Sum, &target.Accumulated, &target.DeadLineDate)
	if err != nil {
		log.Error(err.Error())
		return target, fmt.Errorf("ошибка получения target с id:(%v)", id)
	}

	return target, nil
}

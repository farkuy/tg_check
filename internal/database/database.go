package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Init(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия бд: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к базе данных: %v", err)
	}

	return db, nil
}

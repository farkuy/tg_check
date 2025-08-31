package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

type Query struct {
	name        string
	createQuery string
}

type Storage struct {
	DB *sql.DB
}

var tablesName = []Query{
	{name: "storages", createQuery: creareStoragesTable},
	{name: "storageHistory", createQuery: createStorageHistoryTable},
	{name: "targets", createQuery: creareTargetsTable},
	{name: "targetHistory", createQuery: createTargetHistoryTable},
}

func Init(dbPath string) (*Storage, error) {

	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия бд: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	var wg sync.WaitGroup
	errors := make(chan error)

	wg.Add(len(tablesName))
	for _, table := range tablesName {
		go func(table Query) {
			defer wg.Done()

			checkingTable := checkTable(table.name)

			var exists bool
			if err = db.QueryRow(checkingTable).Scan(&exists); err != nil {
				errors <- fmt.Errorf("ошибка проверки таблицы %v : %v", table.name, err)
			}

			if !exists {
				if _, err = db.Exec(table.createQuery); err != nil {
					errors <- fmt.Errorf("ошибка создания таблицы %v : %v", table.name, err)
				}
			}
		}(table)
	}

	wg.Wait()
	close(errors)

	for range tablesName {
		err = <-errors
		if err != nil {
			return nil, err
		}
	}

	return &Storage{DB: db}, nil
}

func checkTable(tableName string) string {
	return fmt.Sprintf(`
        SELECT EXISTS (
            SELECT 1
            FROM information_schema.tables
            WHERE table_schema = 'public'
              AND table_name = '%s'
        )
    `, tableName)
}

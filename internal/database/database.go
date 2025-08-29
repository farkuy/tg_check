package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Query struct {
	name        string
	createQuery string
}

var tablesName = []Query{
	{name: "storages", createQuery: creareStoragesTable},
	{name: "storageHistory", createQuery: createStorageHistoryTable},
	{name: "targets", createQuery: creareTargetsTable},
	{name: "targetHistory", createQuery: createTargetHistoryTable},
}

func Init(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия бд: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	//Можно попробовать горутину
	start := time.Now()
	for _, table := range tablesName {
		checkingTable := checkTable(table.name)

		var exists bool
		if err = db.QueryRow(checkingTable).Scan(&exists); err != nil {
			return nil, fmt.Errorf("ошибка проверки таблицы %v : %v", table.name, err)
		}

		if !exists {
			if _, err = db.Exec(table.createQuery); err != nil {
				return nil, fmt.Errorf("ошибка создания таблицы: %v", err)
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Время выполнения проверки и создания таблиц: %s\n", elapsed)

	return db, nil
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

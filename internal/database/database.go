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
	{name: "storages", createQuery: createStoragesTable},
	{name: "storageHistory", createQuery: createStorageHistoryTable},
	{name: "targets", createQuery: createTargetsTable},
	{name: "targetHistory", createQuery: createTargetHistoryTable},
}

// Подумать над механизмом миграции
func Init(dbPath string) (*Storage, error) {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия бд: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	var wg sync.WaitGroup
	errorsCh := make(chan error, len(tablesName))

	wg.Add(len(tablesName))
	for _, table := range tablesName {
		go func(table Query) {
			defer wg.Done()

			checkingTable := checkTable(table.name)

			var exists bool
			if err := db.QueryRow(checkingTable).Scan(&exists); err != nil {
				errorsCh <- fmt.Errorf("ошибка проверки таблицы %v : %v", table.name, err)
			}

			if !exists {
				if _, err := db.Exec(table.createQuery); err != nil {
					errorsCh <- fmt.Errorf("ошибка создания таблицы %v : %v", table.name, err)
				}
			}
		}(table)
	}

	wg.Wait()
	close(errorsCh)

	var errorsGroup = []error{}
	for {
		err, ok := <-errorsCh
		if err != nil {
			errorsGroup = append(errorsGroup, err)
		}

		if !ok {
			break
		}
	}

	if len(errorsGroup) > 0 {
		allErros := ""
		for _, val := range errorsGroup {
			allErros = allErros + ", " + val.Error()
		}
		return nil, fmt.Errorf("ошибки инициализации стораджа: %v", allErros)
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

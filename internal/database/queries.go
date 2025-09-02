package database

const (
	createStoragesTable = `CREATE TABLE IF NOT EXISTS storages
	(
    	id SERIAL PRIMARY KEY,
    	sum INT DEFAULT 0,
		accumulated INT DEFAULT 0,
		createDate TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    	deadLineDate TIMESTAMP WITH TIME ZONE
	)`

	createStorageHistoryTable = `CREATE TABLE IF NOT EXISTS storageHistory 
	(
		id SERIAL PRIMARY KEY,
		type TEXT NOT NULL,
		changeSum INT NOT NULL, 
		date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		storageId INT REFERENCES storages(id)    
	)`

	createTargetsTable = `CREATE TABLE IF NOT EXISTS targets
	(
    	id SERIAL PRIMARY KEY,
    	sum INT DEFAULT 0,
		accumulated INT DEFAULT 0,
		storagesId INT REFERENCES storages(id),
		createDate TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    	deadLineDate TIMESTAMP WITH TIME ZONE
	)`

	createTargetHistoryTable = `CREATE TABLE IF NOT EXISTS targetHistory 
	(
		id SERIAL PRIMARY KEY,
		type TEXT NOT NULL,
		changeSum INT NOT NULL, 
		date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		targetId INT REFERENCES targets(id)    
	)
	`
)

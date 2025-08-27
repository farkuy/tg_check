package database

const (
	CreareStoragesTable = `CREATE TABLE IF NOT EXISTS storages
	(
    	id SERIAL PRIMARY KEY,
    	sum INT DEFAULT 0,
		accumulated INT DEFAULT 0,
		createDate TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    	deadLineDate TIMESTAMP WITH TIME ZONE
	)`

	CreareTargetsTable = `CREATE TABLE IF NOT EXISTS targets
	(
    	id SERIAL PRIMARY KEY,
    	sum INT DEFAULT 0,
		accumulated INT DEFAULT 0,
		storagesId INT REFERENCES storages(id),
		createDate TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    	deadLineDate TIMESTAMP WITH TIME ZONE
	)`

	CreateTargetHistoryTable = `CREATE TABLE IF NOT EXISTS targetHistory 
	(
		id SERIAL PRIMARY KEY,
		type TEXT NOT NULL,
		changeSum INT NOT NULL, 
		date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		targetsId INT REFERENCES targets(id)    
	)
	`
)

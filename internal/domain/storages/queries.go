package storages

const (
	postStorageQuery  = "INSERT INTO storages (sum, accumulated, deadLineDate) VALUES ($1, $2, $3) RETURNING *"
	getStorageQuery   = "SELECT * FROM storages WHERE id=$1"
	updateStorageSum  = "UPDATE storages SET sum=$1, accumulated=$2, deadLineDate=$3 WHERE id=$4 RETURNING *"
	deleteStorage     = "DELETE FROM storages WHERE id=$1 RETURNING *"
	checkAddedStorage = "SELECT EXISTS(SELECT 1 FROM storages WHERE id = $1) AS exists_flag;"
)

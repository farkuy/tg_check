package storageHistoty

const (
	postStorageHistoryQuery   = "INSERT INTO storageHistory (type, changeSum, storageId) VALUES ($1, $2, $3) RETURNING *"
	getStorageAllHistoryQuery = "SELECT * FROM storageHistory WHERE storageId IN $1"
	updateStorageHistorySum   = "UPDATE storages SET sum=$1, accumulated=$2, deadLineDate=$3 WHERE id=$4 RETURNING *"
	deleteStorageHistory      = "DELETE FROM storages WHERE id=$1 RETURNING *"
	checkAddedStorageHistory  = "SELECT EXISTS(SELECT 1 FROM storages WHERE id = $1) AS exists_flag;"
)

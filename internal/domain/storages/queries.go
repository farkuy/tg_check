package targets

const (
	postTargetQuery  = "INSERT INTO storages (sum, accumulated, deadLineDate) VALUES ($1, $2, $3) RETURNING *"
	getTargetQuery   = "SELECT * FROM storages WHERE id=$1"
	updateTargetSum  = "UPDATE storages SET sum=$1, accumulated=$2 WHERE id=$3"
	deleteTarget     = "DELETE FROM url WHERE id=$1 RETURNING originalUrl"
	checkAddedTarget = "SELECT EXISTS(SELECT 1 FROM url WHERE originalUrl = $1) AS exists_flag;"
)

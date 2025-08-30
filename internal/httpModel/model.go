package httpModel

type Status string

const (
	StatusOK    Status = "OK"
	StatusError Status = "Error"
)

type Response struct {
	Status Status `json:"status"`
	Error  string `json:"error"`
}

func OK() Response {
	return Response{
		Status: StatusOK,
		Error:  "",
	}
}

func Error(err string) Response {
	return Response{
		Status: StatusError,
		Error:  err,
	}
}

package err

import "net/http"

type Error struct {
	Error string `json:"error"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}

var (
	RespDBDataInsertFailed = []byte(`{"error":"Failed to insert data"}`)
	RespDBDataAccessFailed = []byte(`{"error":"Failed to access data"}`)
	RespDBDataUpdateFailed = []byte(`{"error":"Failed to update data"}`)
	RespDBDataDeleteFailed = []byte(`{"error":"Failed to delete data"}`)

	RespJSONDecodeFailed = []byte(`{"error":"Failed to decode JSON"}`)
	RespJSONEncodeFailed = []byte(`{"error":"Failed to encode JSON"}`)

	RespInvalidURLParamID = []byte(`{"error":"Invalid URL param-ID"}`)
)

func ServerError(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(resp)
}

func BadRequest(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp)
}

func ValidationErrors(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(resp)
}

package crema

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func MakeGenericGetHandler(fn func(map[string]string) (*sql.Rows, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status := GetData(fn, w, r)
		writeResponses(w, data, status)
	}
}

func MakeGenericPostHandler(fn func(*sql.Tx, map[string]string) *sql.Row) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status := PostData(fn, w, r)
		writeResponses(w, data, status)
	}
}

func MakeGenericPutHandler(fn func(*sql.Tx, map[string]string) (sql.Result, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status := PutData(fn, w, r)
		writeResponses(w, data, status)
	}
}

func MakeGenericDeleteHandler(fn func(*sql.Tx, map[string]string) (sql.Result, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status := DeleteData(fn, w, r)
		writeResponses(w, data, status)
	}
}

func writeResponses(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	var response GenericResponse

	if status == http.StatusOK {
		response.Response = GenericHTTPResponse(status)
	} else {
		response.Response = GenericHTTPResponse(status)
	}

	response.Data = data

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	SingleValueQueryString bool
	RawBody                bool
}

var GenericHandler Handler

var requestParameters map[string]string

// MakeGenericGetHandler creates a generic GET-request handler
// TO DO : documentation will be updated soon
func MakeGenericGetHandler(fn func(map[string]string) (*sql.Rows, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status := GetData(fn, w, r)
		writeResponses(w, data, status)
	}
}

// MakeGenericGetHandler creates a generic POST-request handler
// TO DO : documentation will be updated soon
func MakeGenericPostHandler(fn func(*sql.Tx, map[string]string) *sql.Row) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if (*r).Method == "OPTIONS" {
			return
		}
		data, status := PostData(fn, w, r)
		writeResponses(w, data, status)
	}
}

// MakeGenericGetHandler creates a generic PUT-request handler
// TO DO : documentation will be updated soon
func MakeGenericPutHandler(fn func(*sql.Tx, map[string]string) (sql.Result, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if (*r).Method == "OPTIONS" {
			return
		}
		data, status := PutData(fn, w, r)
		writeResponses(w, data, status)
	}
}

// MakeGenericGetHandler creates a generic DELETE-request handler
// TO DO : documentation will be updated soon
func MakeGenericDeleteHandler(fn func(*sql.Tx, map[string]string) (sql.Result, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status := DeleteData(fn, w, r)
		writeResponses(w, data, status)
	}
}

func writeResponses(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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

// DeleteData performs the standard delete mechanism
// TO DO : documentaion will be updated soon
func DeleteData(fn func(*sql.Tx, map[string]string) (sql.Result, error), w http.ResponseWriter, r *http.Request) (interface{}, int) {
	startHandler(r)

	Printf("[HANDLER] started deleting data ...")
	defer Printf("[HANDLER] finished deleting data ...")

	tx, err := BeginTransaction()

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	conditions := requestParameters

	res, err := fn(tx, conditions)

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err.Error(), http.StatusInternalServerError
	}

	Printf(res)
	return fmt.Sprintf("Data has been successfully deleted."), http.StatusOK
}

// PutData performs the standard put mechanism
// TO DO : documentaion will be updated soon
func PutData(fn func(*sql.Tx, map[string]string) (sql.Result, error), w http.ResponseWriter, r *http.Request) (interface{}, int) {
	startHandler(r)

	Printf("[HANDLER] started updating data ...")
	defer Printf("[HANDLER] finished updating data ...")

	tx, err := BeginTransaction()

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	values := requestParameters

	res, err := fn(tx, values)

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err.Error(), http.StatusInternalServerError
	}

	Printf(res)
	return fmt.Sprintf("Data has been successfully updated."), http.StatusOK
}

// PostData performs the standard post mechanism
// TO DO : documentaion will be updated soon
func PostData(fn func(*sql.Tx, map[string]string) *sql.Row, w http.ResponseWriter, r *http.Request) (interface{}, int) {
	startHandler(r)

	Printf("[HANDLER] started inserting data ...")
	defer Printf("[HANDLER] finished inserting data ...")

	tx, err := BeginTransaction()

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	values := requestParameters

	var ID int

	err = fn(tx, values).Scan(&ID)

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err.Error(), http.StatusInternalServerError
	}

	return fmt.Sprintf("Data has been successfully created, id=%v", ID), http.StatusOK
}

// GetData performs the standard get mechanism
// TO DO : documentaion will be updated soon
func GetData(fn func(map[string]string) (*sql.Rows, error), w http.ResponseWriter, r *http.Request) (interface{}, int) {
	startHandler(r)

	Printf("[HANDLER] started getting data ...")
	defer Printf("[HANDLER] finished getting data ...")

	conditions := requestParameters

	rows, err := fn(conditions)

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	data, err := Scan(rows)

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	if data != nil {
		return data, http.StatusOK
	}

	return data, http.StatusNotFound
}

func isGenericHandlerNotYetDefined() bool {
	if !GenericHandler.RawBody && !GenericHandler.SingleValueQueryString {
		return true
	}

	return false
}

func (h *Handler) EnableSingleValueQueryParam() {
	h.SingleValueQueryString = true
}

func (h *Handler) DisableSingleValueQueryParam() {
	h.SingleValueQueryString = false
}

func (h *Handler) EnableRawBody() {
	h.RawBody = true
}

func (h *Handler) DisableRawBody() {
	h.RawBody = false
}

func startHandler(r *http.Request) {
	if isGenericHandlerNotYetDefined() {
		GenericHandler.EnableSingleValueQueryParam()
	}

	requestParameters = make(map[string]string)
	params := mux.Vars(r)

	PopulateParams(params, requestParameters)

	if GenericHandler.SingleValueQueryString {
		PopulateSingleValueQueries(r, requestParameters)
	}

	if GenericHandler.RawBody {
		PopulateRequestBody(r, requestParameters)
	}
}

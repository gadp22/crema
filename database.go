// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func dbConfig() map[string]string {
	conf := ReadDbConfigFile()

	return conf
}

// InitDB sets up db connection based on config file
// TO DO : documentaion will be updated soon
func InitDB() *sql.DB {
	config := dbConfig()

	driver := config["driver"]
	host := config["host"]
	port := config["port"]
	user := config["user"]
	pass := config["pass"]
	dbname := config["dbname"]

	if driver == "psql" {

		Printf("[DB_HELPER] Connecting to PostgreSQL DB ...")

		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, pass, dbname)

		Printf("[DB_HELPER] " + psqlInfo)

		db, err := sql.Open("postgres", psqlInfo)

		HandleError(err)

		err = db.Ping()

		HandleError(err)

		Printf("[DB_HELPER] Successfully connected!")

		return db

	} else if driver == "mysql" {

		connectionString := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname

		db, err := sql.Open("mysql", connectionString)

		HandleError(err)

		return db

	}

	return nil
}

// Scan retrieves data from the sql.Rows dinamically
// TO DO : documentaion will be updated soon
func Scan(rows *sql.Rows) (interface{}, error) {
	config := dbConfig()
	columns, err := rows.Columns()

	var allMaps []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))

		for i := range values {
			pointers[i] = &values[i]
		}

		err = rows.Scan(pointers...)
		resultMap := make(map[string]interface{})

		for i, val := range values {
			if config["driver"] == "mysql" {
				resultMap[columns[i]] = string(val.([]byte))
			} else {
				resultMap[columns[i]] = val
			}
		}

		allMaps = append(allMaps, resultMap)
	}

	return allMaps, err
}

// DeleteData performs the standard delete mechanism
// TO DO : documentaion will be updated soon
func DeleteData(fn func(*sql.Tx, map[string]string) (sql.Result, error), w http.ResponseWriter, r *http.Request) (interface{}, int) {
	Printf("[DATABASE_HELPER] started deleting data ...")
	defer Printf("[DATABASE_HELPER] finished deleting data ...")

	tx, err := BeginTransaction()

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	var conditions map[string]string
	conditions = make(map[string]string)

	params := mux.Vars(r)

	conditions = PopulateSingleValueQueries(r, conditions)
	conditions = PopulateParams(params, conditions)

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
	Printf("[DATABASE_HELPER] started updating data ...")
	defer Printf("[DATABASE_HELPER] finished updating data ...")

	tx, err := BeginTransaction()

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	var values map[string]string
	values = make(map[string]string)

	params := mux.Vars(r)

	values = PopulateSingleValueQueries(r, values)
	values = PopulateParams(params, values)

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
	Printf("[DATABASE_HELPER] started inserting data ...")
	defer Printf("[DATABASE_HELPER] finished inserting data ...")

	tx, err := BeginTransaction()

	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	var values map[string]string
	values = make(map[string]string)

	params := mux.Vars(r)

	values = PopulateSingleValueQueries(r, values)
	values = PopulateParams(params, values)

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
	Printf("[DATABASE_HELPER] started getting data ...")
	defer Printf("[DATABASE_HELPER] finished getting data ...")

	var conditions map[string]string
	conditions = make(map[string]string)

	params := mux.Vars(r)

	conditions = PopulateSingleValueQueries(r, conditions)
	conditions = PopulateParams(params, conditions)

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

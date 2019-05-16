// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

		Printf("[DATABASE] Connecting to PostgreSQL DB ...")

		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, pass, dbname)

		Printf("[DATABASE] " + psqlInfo)

		db, err := sql.Open("postgres", psqlInfo)

		HandleError(err)

		err = db.Ping()

		HandleError(err)

		Printf("[DATABASE] Successfully connected!")

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

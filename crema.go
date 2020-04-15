// Copyright 2019 The Crema Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crema

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

// InitServer initializes the core system by setting up
// the database connection, routing services and system logger
func InitServer() *Server {
	initLogger()

	LogPrintf("[MAIN] Initializing Server ...")

	db := initDatabase()
	router := initRoutes()

	return &Server{db, router}
}

func initDatabase() *sql.DB {
	LogPrintf("[MAIN] Initializing Database ...")

	db := InitDB()

	if db != nil {
		SetDB(db)

		return db
	} else {
		return nil
	}

}

func initLogger() {
	log.Println("[MAIN] Initializing Logger ...")

	InitLogFiles()
	Printf("[MAIN] Initializing Logger ...")
}

func initRoutes() *mux.Router {
	LogPrintf("[MAIN] Initializing Endpoints ...")

	router := mux.NewRouter()

	return router
}

// AddRoutes is used to create HTTP request routing
// TO DO: documentation to be updated soon
func (s *Server) AddRoutes(method string, routes string, handler func(http.ResponseWriter, *http.Request)) {
	s.Router.HandleFunc(routes, handler).Methods(method)
}

// AddRoutesOp with Options
func (s *Server) AddRoutesOp(method string, routes string, handler func(http.ResponseWriter, *http.Request)) {
	s.Router.HandleFunc(routes, handler).Methods(method, "OPTIONS")
}

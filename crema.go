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

<<<<<<< HEAD
func InitServer() *Server {
	InitLogger()
=======
// InitServer initializes the core system by setting up
// the database connection, routing services and system logger
func InitServer() *server {
	initLogger()
>>>>>>> 50ebc3a... [refactor][add] renamed source files, removed unnecessary exported functions, added documentations

	LogPrintf("[MAIN] Initializing Server ...")

	db := initDatabase()
	router := initRoutes()

	return &Server{db, router}
}

func initDatabase() *sql.DB {
	LogPrintf("[MAIN] Initializing Database ...")

	db := InitDB()
	SetDB(db)

	return db
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

<<<<<<< HEAD
func (s *Server) AddRoutes(method string, routes string, handler func(http.ResponseWriter, *http.Request)) {
=======
// AddRoutes is used to create HTTP request routing
// TO DO: documentation to be updated soon
func (s *server) AddRoutes(method string, routes string, handler func(http.ResponseWriter, *http.Request)) {
>>>>>>> 50ebc3a... [refactor][add] renamed source files, removed unnecessary exported functions, added documentations
	s.Router.HandleFunc(routes, handler).Methods(method)
}

package main

import (
	"log"
	"logger"
	"net/http"

	"crema"
)

func main() {
	crema.InitLogger()

	server := crema.InitServer()

	server.AddRoutes(http.MethodGet, "/hello", hello)

	logger.LogPrintf("[MAIN] Server is running, listening to port 8001 ....")
	log.Fatal(http.ListenAndServe(":8001", server.Router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

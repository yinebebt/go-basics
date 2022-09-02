package main

import (
	"gotest/BuildApp/server"
	"log"
	"net/http"
)

func main() {
	server:= &(server.PlayerServer{})
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5050", handler))
}

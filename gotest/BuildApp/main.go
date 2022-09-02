package main

import (
	"gotest/BuildApp/server"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5050", handler))
}

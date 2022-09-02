package main

import (
	"gotest/BuildApp/server"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	// server:= &(server.PlayerServer{})
	server := &(server.PlayerServer{Store: &InMemoryPlayerStore{}})
	// handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5050", server)) //you can pass hander too
}

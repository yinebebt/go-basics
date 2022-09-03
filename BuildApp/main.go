package main

import (
	"gotest/BuildApp/server"
	"log"
	"net/http"
)

func main() {
	// server:= &(server.PlayerServer{})
	server := &(server.PlayerServer{Store: &InMemoryPlayerStore{}})
	// handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5050", server)) //you can pass hander too
}

// in_memoryPLayeStore.go
type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}  //used to initialize in other place

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

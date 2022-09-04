package main

import (
	"log"
	"net/http"
)

// I created this var to have a global pointer to PLaySrver in testing and main, But lately get it less important.
// var pstore = &PlayerServer{NewInMemoryPlayerStore()}

func main() {
	// server := &PlayerServer{NewInMemoryPlayerStore()} // the reason here we use pointer is since *PlayerStore does implement serverHTTP
	server := NewPlayerServer(NewInMemoryPlayerStore())

	// handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", server)) //you can pass hander too
	// fmt.Println("pointer to PlayServer: ", pstore)
}

// func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
// 	return 123
// }

// in_memoryPLayeStore.go
type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
} //used to initialize in other place

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	score := i.store[name]
	return score
}

func (i *InMemoryPlayerStore) GetLeague() []Player{
	return nil
}

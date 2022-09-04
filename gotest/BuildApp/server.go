package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// server.go
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(string)
	GetLeague() []Player
}

type PlayerServer struct {
	Store        PlayerStore
	http.Handler // we can use  router *http.ServeMux too
}

// creating JSON data Model
type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.Store = store
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json") //setting header content-type as a applicatio/json type
	json.NewEncoder(w).Encode(p.Store.GetLeague())
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

//If we were using router *Server.Mux
// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	p.router.ServeHTTP(w, r)
// 	}// to handle requests : pattern with handler

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// getRequest is a helper function
func getScoreRequest(playerName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", playerName), nil)
	return req
}

func PostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// func (p *PlayerServer) getLeagueTable() []Player {
// 	return  []Player{
// 		{"Bereket", 32},
// 		{"Abdi", 20},
// 		{"Israiel", 14},
// 	}
// }

// server_integration_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	// server := PlayerServer{NewInMemoryPlayerStore()}  //initialize PLayeSore with a map with default value
	server := NewPlayerServer(NewInMemoryPlayerStore())

	// server := pstore  // If you want to use the global pointer, but it seems meaning less testing should run independently.
	player := "Abel"
	server.ServeHTTP(httptest.NewRecorder(), PostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), PostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), PostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, getScoreRequest(player))
	asserStatus(t, response.Code, http.StatusOK)

	assertScore(t, response.Body.String(), "3") //since we called PostNewRequest 3 times above, via InMemoryStore

	//Below test was aim to see the pointer address where PlayStore points to.
	// t.Run("seeing pointer to  PlayServer", func(t *testing.T) {
	// 	a := server
	// 	if !reflect.DeepEqual(a, "&{0xc000014028}") {
	// 		t.Errorf("got %v want %v", a, "&{0xc000014028}")
	// 	}
	// })
}

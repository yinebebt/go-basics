// server_integration_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Abel"
	server.ServeHTTP(httptest.NewRecorder(), PostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), PostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), PostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, getScoreRequest(player))
	asserStatus(t, response.Code, http.StatusOK)

	assertScore(t, response.Body.String(), "3") //since we called POstNewRequest 3 times above, via InMemoryStore
}

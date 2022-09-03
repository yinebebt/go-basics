// server_integration_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	server "gotest/BuildApp/server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()  
	servers := server.PlayerServer{store}
	player := "Pepper"


	servers.ServeHTTP(httptest.NewRecorder(), server. (player))
	servers.ServeHTTP(httptest.NewRecorder(), server.PostWinRequest(player))
	servers.ServeHTTP(httptest.NewRecorder(), server.PostWinRequest(player))

	response := httptest.NewRecorder()
	servers.ServeHTTP(response, server.getScoreRequest(player))
	server.asserStatus(t, response.Code, http.StatusOK)

	servers.asserResponseBody(t, response.Body.String(), "3")
}

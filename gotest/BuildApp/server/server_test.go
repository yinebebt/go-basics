package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Abel's score", func(t *testing.T) {
		request := getScoreRequest("Abel")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		assertScore(t, got, "20")
	})
	t.Run("returns Tesfa's score", func(t *testing.T) {
		request := getScoreRequest("Tesfa")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		assertScore(t, got, "10")
	})
}

// getRequest is a helper function
func getScoreRequest(playerName string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", playerName), nil)
	return req
}

// assertScore is a helper function
func assertScore(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

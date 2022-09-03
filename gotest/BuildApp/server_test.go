package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Abel":  20,
			"Tesfa": 10,
		}, nil,
	}
	server := &PlayerServer{&store}

	t.Run("returns Abel's score", func(t *testing.T) {
		request := getScoreRequest("Abel")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got_body := response.Body.String()
		got_status := response.Code

		assertScore(t, got_body, "20")
		asserStatus(t, got_status, 200)
	})
	t.Run("returns Tesfa's score", func(t *testing.T) {
		request := getScoreRequest("Tesfa")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got_body := response.Body.String()
		got_status := response.Code

		assertScore(t, got_body, "10")
		asserStatus(t, got_status, 200)
	})
	//server_test.go
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := getScoreRequest("Meklit")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got_body := response.Body.String()
		got_status := response.Code

		assertScore(t, got_body, "0")
		asserStatus(t, got_status, 404)

	})
}

// assertScore is a helper function
func assertScore(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func asserStatus(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

//for POST, to update winner's score, the above were for Get

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Abel"
		request := PostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		asserStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
	})

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Tesfa"
		request := PostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		asserStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 2 { //2 because we run two subtests and RecordWin called two times.
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 2)
		}

		if store.winCalls[1] != "Tesfa" {
			t.Errorf("got %s calls to RecordWin want %s", store.winCalls[1], "Tesfa")
		}
	})
}

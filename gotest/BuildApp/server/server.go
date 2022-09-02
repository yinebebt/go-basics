package server

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer has a signature similar with HandlerFunc
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(name string) string {
	if name == "Abel" {
		return "20"
	}

	if name == "Tesfa" {
		return "10"
	}
	return "0"
}

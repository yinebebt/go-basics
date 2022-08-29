package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebRacer(t *testing.T) {
	slowServer := serverdelay(20 * time.Millisecond)
	fastServer := serverdelay(0)
	//Delay a few duration to have a difference  in the response, it is just testing trial/learning purpose

	slowURL := slowServer.URL //URL is a filed to the server struct
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	slowServer.Close() //close the mocking/ httptest server after use.
	fastServer.Close()
}

func serverdelay(t time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t)
		w.WriteHeader(http.StatusOK)
	})) //NewServer takes anonymous function
}

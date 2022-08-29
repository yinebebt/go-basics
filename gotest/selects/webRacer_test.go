package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// //v-1
func TestWebRacer(t *testing.T) {
	slowServer := makeServerDelayed(20 * time.Millisecond)
	fastServer := makeServerDelayed(0)
	//Delay a few duration to have a difference  in the response, it is just testing trial/learning purpose

	slowURL := slowServer.URL //URL is a filed to the server struct
	fastURL := fastServer.URL

	want := fastURL
	got, _ := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	slowServer.Close() //close the mocking/ httptest server after use.
	fastServer.Close()

	//v-2
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeServerDelayed(11 * time.Second)
		serverB := makeServerDelayed(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one", err)
		}
		// if err != nil {
		// 	t.Error(err)
		// }
	})
}

func makeServerDelayed(t time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t)
		w.WriteHeader(http.StatusOK)
	})) //NewServer takes anonymous function
}

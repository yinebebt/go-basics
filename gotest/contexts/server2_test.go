package contexts

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore2 struct {
	response string
}

func (s *SpyStore2) fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1) //channel of buffer size 1
	go func() {
		var result string
		for _, c := range s.response { //?range loop through character/rune of s.response
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return // leave the for loop
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		} //we accept and store string char by char, in a step ways- slowly
		data <- result //place the result string into the channel data
	}()

	select {
	case <-ctx.Done(): //if it is not yet done(?)
		return "", ctx.Err() //If Done is not yet closed, Err returns nil. If Done is closed, Err returns a non-nil error explaining why
	case res := <-data:
		return res, nil
	} //wait for goroutine to finish its work or for the cancellation to occur.
}

//inorder not to write response on error
//roll your own spy for http.ResponseWriter (if you need it)

type SpyResponseWriter struct {
	written bool
}

//SpyResponseWriter implements http.ResponseWriter

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}
func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer2(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore2{response: data}
		svr := Server2(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}
		svr.ServeHTTP(response, request)
		if response.written { //meaning once we cancel the request, it shouldn't wrote to the respnsewritter- the spy in this case,
			t.Error("a response should not have been written")
		} //not getting this condition means, things are going as we expect; no write happens on error!
	})

}

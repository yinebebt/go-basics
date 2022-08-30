package contexts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct { // a stub for Store
	response string
}
type SpyStore struct {
	response  string
	cancelled bool
}

// func (s *StubStore) fetch() string {
// 	return s.response
// }

// v-2
func (s *SpyStore) fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) pcancelT() { // the reason I make the name like this is to assert that the name cancel have no relation with the contetxwith cancel canceFUnc.
	//this method is created ust to spy/see either the context was cancelled within the defined time- in AfterFUnc
	s.cancelled = true
}

func TestServer(t *testing.T) {
	// data := "hello, world"
	// svr := Server(&StubStore{response:data})
	// request := httptest.NewRequest(http.MethodGet, "/", nil) //const http.MethodGet untyped string = "GET"
	// response := httptest.NewRecorder()                       //NewRecorder returns an initialized ResponseRecorder.
	// svr.ServeHTTP(response, request)
	// if response.Body.String() != data {
	// 	t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	// }

	// v-2
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		// fmt.Println(cancel)

		request = request.WithContext(cancellingCtx) // send request with new context with a cancel after
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		if !store.cancelled { //meaning "if the method `cancel` is not called...
			t.Error("store was not told to cancel")
		}
	})

	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
		if store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})
	// g,j:= context.CancelFunc()
}

package contexts

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	fetch() string
	pcancelT()
}

type Store2 interface {
	fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// store.cancel()

		// fmt.Fprint(w, store.fetch())
		//v-2
		ctx := r.Context()
		data := make(chan string, 1)
		go func() {
			data <- store.fetch()
		}()
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.pcancelT()
		}
	}
}

func Server2(store Store2) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.fetch(r.Context()) //data = result
		if err != nil {
			return // todo: log error however you like
		} // to avoid writing a  response on errors other than nil.
		fmt.Fprint(w, data)
	}
}

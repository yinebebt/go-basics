package selects

import (
	"fmt"
	"net/http"
	"time"
)

//v-1
// func Racer(a, b string) (winner string) {
// 	aDuration := responseTime(a)

// 	bDuration := responseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}
// 	return b
// }

// func responseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url) //Get sends a request and return a response
// 	return time.Since(start)
// }

// v-2
func Racer(a, b string) (winner string, error error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
} //time.After returns a chan (like ping) and will send a signal down it after the amount of time you define.

func ping(a string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(a)
		close(ch)
	}()
	return ch
}

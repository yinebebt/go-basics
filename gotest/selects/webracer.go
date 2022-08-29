package selects

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := timeDelay(a)

	bDuration := timeDelay(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func timeDelay(url string) time.Duration {
	start := time.Now()
	http.Get(url) //Get sends a request and return a response
	return time.Since(start)
}

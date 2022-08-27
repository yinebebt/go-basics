package concurrency

type webchecker func(url string) bool

type result struct {
	string //leave the field name empty, if you doute what to name it
	bool
}

func checkWebsites(w webchecker, urls []string) map[string]bool {
	results := make(map[string]bool)   //used to store returned values
	resultChannel := make(chan result) //creating a channel of type result.

	for _, url := range urls {
		go func(a string) {
			// result[a] = w(a)
			resultChannel <- result{a, w(a)}
		}(url) //loopclosure issue solved!
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	// time.Sleep(1 * time.Second) //since the for loop fast and did't give a time to the go routin,
	return results
}

// 21337796  ~ 0.02s  - via go routines
//2034683179 ~ 2.s - without goroutines

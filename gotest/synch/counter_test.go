package synch

import (
	"sync"
	"testing"
)

func TestCOunter(t *testing.T) {
	t.Run("counter test with 3 increment", func(t *testing.T) {
		counter := counter{}
		// counter.Inc()
		// counter.Inc()
		// counter.Inc()

		//v-2
		countdown := 1000
		var wg sync.WaitGroup
		wg.Add(countdown)

		for i := 0; i < countdown; i++ {
			go func() {
				counter.Inc()
				wg.Done() //Done decrements the WaitGroup counter by one.
			}()
		}
		wg.Wait() //Wait blocks until the WaitGroup counter is zero.
		//each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.
		assertCounter(t, &counter, 1000)
	})

}

func assertCounter(t testing.TB, got *counter, want int) { //make the pointer fix the issue where mutex loc is copied by assertCouter.
	if got.value != want {
		t.Errorf("Got %d want %d", got, want)
	}
}

// or you can use NewCOunter, once defining a the constructor in your test to initialize the counter
func NewCounter() *counter {
	return &counter{}
}

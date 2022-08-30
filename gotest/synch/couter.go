package synch

import "sync"

type counter struct {
	//let I add a mutex to kep our counter stable even with several go routines
	mu    sync.Mutex //A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	value int
}

func (c *counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *counter) val() int {
	return c.value
}

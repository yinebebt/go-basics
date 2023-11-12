package main

import (
	"fmt"
	"sort"
	"sync"
)

// go's map doesn't support concurrent read/r=write, and
//there is no definition if a routine do that onto a map at a time
//one way to do is using sync.RWMutex

var counter = struct {
	sync.RWMutex
	m map[string]int
}{m: make(map[string]int)}

func main() {
	for i := 0; i < 6; i++ {
		counter.m[fmt.Sprintf("key%d", i)] = i * 10
	}
	//To read from the counter, take the read lock; (inside some goroutine)
	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)

	//To write to the counter, take the-write lock;
	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

	//map entries are not ordered in range iterations, to do this you need to declare a separate data structure like
	var m map[int]string
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println("from cache", foo)
	}

	// Since Go is statically typed, and cache values can be anything, type
	// assertion is needed when values are being passed to functions that don't
	// take arbitrary types, (i.e. interface{}).
	foo, found = c.Get("foo")
	if found {
		varFoo := foo.(string) //type casting
		fmt.Println("string of foo", varFoo)
	}

	// Want performance? Store pointers!
	mystruct := &MyStruct{}
	c.Set("foo", mystruct, cache.DefaultExpiration)
	if x, found := c.Get("foo"); found {
		foo = x.(*MyStruct)
		fmt.Println("from pointer cache", foo)
	}
}

type MyStruct struct {
	name string
}

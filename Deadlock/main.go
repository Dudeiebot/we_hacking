package main

import (
	"appengine/memcache"
)

func main() {

	type dude struct {
		heisadude string
		notadude  string
		err       error
	}
	dataChan := make(chan dude) //created my data type here and added it to my channel

	dataChan <- dude{}
	go func() {

	}()

	go func() {

	}()

	go func() {

	}()

	if Err != memcache.ErrCacheMiss {
		c.Errorf("memcache get %q: %v", ID, err)
	}
}

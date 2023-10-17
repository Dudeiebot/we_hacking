package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// data := <- a read from channel a
// a <- data  write to channel a

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(50)
}

// Like maps and slices, channels must be created before use, that is where the make chan comes in.
func main() {
	dataChan := make(chan int) // Channel can be buffered "make (chan int, 3)"

	go func() { // a  new go routine cause by default a channel doesnot have space to store data or maybe you add the require space storage to dataChan i.e dataChan := make(chan int 2)
		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				// defer wg.Done()  so these is just added when my wg.add(2) and we need to tell it we are done twice and these are just for testing sake and getting thing safe and done
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Printf("n : %d\n", n)
	}
}

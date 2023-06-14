package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(50)
}

func main() {
	dataChan := make(chan int)

	go func() { // a  buffered go routine or maybe you add the require output to dataChan i.e dataChan := make(chan int 2)
		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				//defer wg.Done()  so these is just added when my wg.add(2) and we need to tell it we are done twice and these are just for testing sake and getting thing safe and done
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

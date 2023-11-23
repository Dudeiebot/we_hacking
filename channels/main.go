package main

import (
	"fmt"
)

// import (
//   "math/rand"
//   "sync"
//    "time"
// )

/*

First: you create your channel with the make function and it can be buffered or unbuffered
data := <- a read/receive from channel and a
a <- data  write/send to channel a
and finally clode your channel anyway you like to close it

*/

// func DoWork() int {
// 	time.Sleep(time.Second)
// 	return rand.Intn(50)
// }
//
// // Like maps and slices, channels must be created before use, that is where the make chan comes in.
// func main() {
// 	dataChan := make(chan int) // Channel can be buffered "make (chan int, 3)"
//
// 	go func() { // a  new go routine cause by default a channel doesnot have space to store data or maybe you add the require space storage to dataChan i.e dataChan := make(chan int 2)
// 		wg := &sync.WaitGroup{}
// 		for i := 0; i < 10; i++ {
// 			wg.Add(1)
// 			go func() {
// 				defer wg.Done()
// 				// defer wg.Done()  so these is just added when my wg.add(2) and we need to tell it we are done twice and these are just for testing sake and getting thing safe and done
// 				result := DoWork()
// 				dataChan <- result
// 			}()
// 		}
// 		wg.Wait()
// 		close(dataChan)
// 	}()
//
// 	for n := range dataChan {
// 		fmt.Printf("n : %d\n", n)
// 	}
// }

//These is a new func showing buffered channels
// func write(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("successfully wrote", i, "to ch")
// 	}
// 	close(ch)
// }
//
// func main() {
// 	ch := make(
// 		chan int,
// 		2,
// 	) // remember buffered channel i am using it here and it is just outputting 2 input from func write before it continue with the remaining
// 	go write(ch)
// 	time.Sleep(2 * time.Second)
// 	for v := range ch {
// 		fmt.Println("read value", v, "from ch")
// 		time.Sleep(2 * time.Second)
// 	}
// }

func main() {
	datach := make(chan string, 3)
	datach <- "hello"
	datach <- "world"
	close(datach)
	fmt.Println("These is the capacity of datach:", cap(datach))
	fmt.Println("These is the length of datach:", len(datach))
	fmt.Println("We are reading a value from datach:", <-datach)
	fmt.Println("Here is our new lenght after we read that value:", len(datach))
}

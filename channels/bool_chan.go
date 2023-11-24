package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	resultChan := make(chan bool)

	go checkRandomNumber(resultChan)

	result := <-resultChan
	fmt.Println(result)
}

func checkRandomNumber(resultChan chan bool) {
	// Simulating some asynchronous work
	time.Sleep(2 * time.Second)

	// Generating a random number between 0 and 1
	randomNumber := rand.Float64()

	// Sending the result of the condition through the channel
	resultChan <- randomNumber > 0.5
	close(resultChan)
}

// can print out true or false randomly as it likes

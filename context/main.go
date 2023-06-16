package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your timeframe (in milliseconds): ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input) //trim leading and trailing whitespace

	timer, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	t
	timeoutDuration := time.Millisecond * time.Duration(timer)
	timeoutContext, cancel := context.WithTimeout(context.Background(), timeoutDuration) //understands your timeout well before pushing
	defer cancel()
	//create http requests
	req, err := http.NewRequestWithContext(timeoutContext, "GET", "https://via.placeholder.com/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	//performs http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	imagedata, _ := io.ReadAll(res.Body)

	fmt.Printf("downloaded image size %d\n", len(imagedata))

}

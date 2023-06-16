package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your timeframe (in second): ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input) //trim leading and trailing whitespace

	timer, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {

		timeoutDuration := time.Second * time.Duration(timer)
		timeoutContext, cancel := context.WithTimeout(c.Request.Context(), timeoutDuration)
		defer cancel()

		req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://yahoo.com", nil)
		if err != nil {
			panic(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		c.Data(200, "text/html", data)
	})
	r.Run()
}

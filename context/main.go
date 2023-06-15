package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	//create http requests
	req, err := http.NewRequest("GET", "https://via.placeholder.com/2000x2000", nil)
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

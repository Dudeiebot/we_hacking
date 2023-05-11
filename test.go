package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("server failed to start %v", err)
	}
}

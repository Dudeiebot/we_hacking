package main

import (
	"fmt"

	"rsc.io/quote"
)

// Returning some certain quotes in golang, not something serious then here comes quote.Hello() used for those shii
func main() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
}

// func hellos(names []string) (map[string]string, error) {
// 	messages := make(map[string]string)
// 	for _, name := range names {
// 		message, err := hello(name)
// 		if err != nil {
// 			return err, nil
// 		}
// 		messages[name] = message
// 	}
// 	return messages, nil
// }

// so there is quote.Hello() for returning greeting, not all that necessary tho

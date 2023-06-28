package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
}

func hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := hello(name)
		if err != nil {
			return err, nil
		}
		messages[name] = message
	}
	return messages, nil
}

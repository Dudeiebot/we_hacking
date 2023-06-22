package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Add[T constraints.Ordered](a T, b T) T {
	return a + b
}

func main() {
	result := Add(1.1, 2.2)
	fmt.Printf("result: %+v\n", result)
}

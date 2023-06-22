package main

import (
	"fmt"
)

type UserID float64

func Add[T int | ~float64](a T, b T) T {
	return a + b
}

func main() {
	a := UserID(1)
	b := UserID(2)
	result := Add(a, b)
	fmt.Printf("result: %+v\n", result)
}

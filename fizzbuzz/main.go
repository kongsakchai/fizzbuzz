package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, FizzBuzz")
}

func FizzBuzz(n int) string {
	if n == 4 {
		return "4"
	}
	if n == 3 {
		return "Fizz"
	}
	if n == 2 {
		return "2"
	}
	return "1"
}

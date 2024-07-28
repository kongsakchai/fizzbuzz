package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, FizzBuzz")
}

func FizzBuzz(n int) string {
	if n == 5 {
		return "Buzz"
	}
	if n == 3 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}

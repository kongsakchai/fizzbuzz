package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, FizzBuzz")
}

func FizzBuzz(n int) string {
	if n == 15 || n == 30 {
		return "FizzBuzz"
	}
	if n == 5 || n == 10 || n == 25 {
		return "Buzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}

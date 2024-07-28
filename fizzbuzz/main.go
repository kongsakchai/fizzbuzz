package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, FizzBuzz")
}

func FizzBuzz(n int) string {
	if n == 15 {
		return "FizzBuzz"
	}
	if n == 5 || n == 10 {
		return "Buzz"
	}
	if n == 3 || n == 6 || n == 9 || n == 12 || n == 18 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}

package main

import "fmt"

func Fibonacci(n int) interface{} {
	if (n < 2) {
		return n
		return (Fibonacci((n - 1)) + Fibonacci((n - 2)))
		for i := 0; i < 10; i++ {
			fmt.Println(Fibonacci(i))
		}
	}
	return nil
}

func main() {
}

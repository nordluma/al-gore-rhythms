package main

import "fmt"

func testFib(msg string, n int, fn func(int) int) {
	fmt.Println(msg)
	fmt.Println(fn(n))
	fmt.Println("==================")
}

func main() {
	testFib("Normal Fibonacci", 40, fibonacci)
	testFib("Memoized Fibonacci", 40, memFibonacci)
}

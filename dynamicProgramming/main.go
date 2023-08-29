package main

import "fmt"

func testFib(msg string, n int, fn func(int) int) {
	fmt.Println(msg)
	fmt.Println(fn(n))
	fmt.Println("==================")
}

func main() {
	testFib("Normal Fibonacci", 50, fibonacci)
	testFib("Bottom up Fibonacci", 50, btmUpFibonacci)
	testFib("Memoized Fibonacci", 50, memFibonacci)
	testFib("Bottom up memoized Fibonacci", 50, btmUpMemFibonacci)
}

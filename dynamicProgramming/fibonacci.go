package main

func fibonacci(n int) int {
	if n == 1 || n == 0 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func _memFibonacci(n int, cache map[int]int) int {
	if n == 1 || n == 0 {
		return n
	}

	if _, ok := cache[n]; !ok {
		cache[n] = _memFibonacci(n-1, cache) + _memFibonacci(n-2, cache)
	}

	return cache[n]
}

func memFibonacci(n int) int {
	cache := make(map[int]int)
	return _memFibonacci(n, cache)
}

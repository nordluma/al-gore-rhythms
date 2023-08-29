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

func btmUpMemFibonacci(n int) int {
	if n == 1 || n == 0 {
		return 0
	}

	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1

	for i := 2; i < n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n-1] + cache[n-2]
}

func btmUpFibonacci(n int) int {
	if n == 0 {
		return 0
	}

	a := 0
	b := 1
	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c
	}

	return a + b
}

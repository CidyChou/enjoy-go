package util

func Fibonacci(n int) int {

	if n == 0 {
		return -1
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

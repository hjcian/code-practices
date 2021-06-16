package fibonaccinumber

func fib(n int) int {
	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Fibonacci Number.
	// Memory Usage: 1.9 MB, less than 32.83% of Go online submissions for Fibonacci Number.
	if n <= 1 {
		// base cases
		return n
	}
	fn_sub_2 := 0
	fn_sub_1 := 1
	fn := 0
	for n > 1 {
		fn = fn_sub_2 + fn_sub_1
		fn_sub_2 = fn_sub_1
		fn_sub_1 = fn
		n--
	}
	return fn
}

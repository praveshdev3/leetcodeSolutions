func fib(n int) int {
	f0 := 0
	f1 := 1
	if n == 0 {
		return f0
	}
	if n == 1 {
		return f1
	}
	result := f0 + f1
	for i := 2; i < n; i++ {
		f0 = f1
		f1 = result
		result = f0 + f1
	}
	return result
}

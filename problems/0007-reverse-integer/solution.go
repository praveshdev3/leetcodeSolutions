func reverse(x int) int {
	result := 0
	for x != 0 {
		num := x % 10
		x = x / 10
		result = result*10 + num
		if result > 2147483647 || result < -2147483648 {
			return 0
		}
	}
	return result
}

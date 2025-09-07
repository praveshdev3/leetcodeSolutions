func beautySum(s string) int {
	n := len(s)
	var totalBeauty int
	for i := range s {
		var charFreq [26]int
		countPerFreq := make([]int, 500)
		max := 0
		min := math.MaxInt32
		for j := i; j < n; j++ {
			ch := s[j] - 'a'
			f := charFreq[ch]
			countPerFreq[f]--
			countPerFreq[f+1]++
			charFreq[ch]++
			if f+1 > max {
				max = f + 1
			}
			if f+1 < min || (min == f && countPerFreq[min] == 0) {
				min = f + 1
			}
			if max > min {
				totalBeauty += max - min
			}
		}
	}
	return totalBeauty
}

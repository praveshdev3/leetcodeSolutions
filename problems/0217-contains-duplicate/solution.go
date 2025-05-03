func containsDuplicate(nums []int) bool {
	mp := make(map[int]int)
	for _, num := range nums {
		_, ok := mp[num]
		if !ok {
			mp[num] = 1
		} else {
			return true
		}
	}
    return false
}

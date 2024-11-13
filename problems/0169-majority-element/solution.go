func majorityElement(nums []int) int {
	count := 1
	majorityElement := nums[0]

	for i := 1; i < len(nums); i++ {
		if count <= 0 {
			majorityElement = nums[i]
			count = 1
		} else {
			if majorityElement == nums[i] {
				count++
			} else {
				count--
			}
		}
	}
	return majorityElement
}

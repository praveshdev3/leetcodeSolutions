func sortColors(nums []int) {
	zeroIdx := 0
	twoIdx := len(nums) - 1

	for i := 0; i <= twoIdx; i++ {
		if nums[i] == 0 {
			nums[i], nums[zeroIdx] = nums[zeroIdx], nums[i]
			zeroIdx++
		} else if nums[i] == 2 {
			nums[i], nums[twoIdx] = nums[twoIdx], nums[i]
			twoIdx--
            i--
		}
	}
}

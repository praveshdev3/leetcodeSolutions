func sumIndicesWithKSetBits(nums []int, k int) int {
    sum := 0
    for i := range len(nums) {
        if bits.OnesCount(uint(i)) == k {
            sum += nums[i]
        }
    }

    return sum
}

func countPairs(nums []int, target int) int {
    sort.Ints(nums)

    left, right, ans := 0 , len(nums)-1, 0
    for left < right {
        if nums[left] + nums[right] < target {
            ans = ans + (right-left)
            left++
        }else {
            right--
        }
    }

    return ans
}

func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)

    windowStart := 0
    result := 1
    sum := 0

    for windowEnd := 0; windowEnd < len(nums); windowEnd++{
        sum += nums[windowEnd]
        if nums[windowEnd] * (windowEnd - windowStart + 1) > sum + k{
            sum -= nums[windowStart]
            windowStart++
        }
        result = max(result, windowEnd-windowStart+1)
    }

    return result
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

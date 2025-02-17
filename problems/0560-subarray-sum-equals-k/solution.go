func subarraySum(nums []int, k int) int {
    prefix := make(map[int]int)
    prefix[0] = 1
    sum := 0
    count := 0

    for _,num := range nums{
        sum += num

        count += prefix[sum-k]
        prefix[sum]++
    }

    return count
}

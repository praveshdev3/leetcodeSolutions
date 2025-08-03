func subarraySum(nums []int, k int) int {
    prefix := make(map[int]int)
    sum := 0
    count := 0
    prefix[0]++

    for _,num := range nums{
        sum += num

        count += prefix[sum-k]
        prefix[sum]++
    }

    return count
}

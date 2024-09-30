func minOperations(nums []int, k int) int {
    minOperations := 0
    for _,num := range nums{
        if num < k {
            minOperations++
        }
    }
    return minOperations
}

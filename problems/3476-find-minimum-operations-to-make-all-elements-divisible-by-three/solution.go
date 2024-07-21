func minimumOperations(nums []int) int {
    count := 0
    for _, num := range nums {
        if num % 3 != 0 {
            count++
        }
    }
    return count
}

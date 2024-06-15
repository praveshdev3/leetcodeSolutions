func twoSum(nums []int, target int) []int {
    mapForNums := make(map[int]int)
    for i, num := range nums{
        complement := target - num
        if j, ok := mapForNums[complement]; ok {
            return []int{i,j}
        }
        mapForNums[num] = i
    }
    return nil
}

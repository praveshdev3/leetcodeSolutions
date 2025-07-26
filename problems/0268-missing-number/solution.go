func missingNumber(nums []int) int {
    sum,arrSum := 0,0
    for i,num := range nums{
        sum += i+1
        arrSum += num
    }
    return sum-arrSum
}

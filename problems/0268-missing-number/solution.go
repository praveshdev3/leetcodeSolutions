func missingNumber(nums []int) int {
    sum,arrSum := 0,0

    for i:=0; i<len(nums); i++{
        sum = sum+i+1
        arrSum += nums[i] 
    }

    return sum-arrSum
}

func maxAscendingSum(nums []int) int {
    maxSum := 0
    currSum := 0
    for i:=0; i<len(nums); i++{
        if i==0 || nums[i] > nums[i-1]{
            currSum += nums[i]
            if currSum > maxSum{
                maxSum = currSum
            }
        } else{
            currSum = nums[i]
        }
    }
    return maxSum
}

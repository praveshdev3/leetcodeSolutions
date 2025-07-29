func maxSubArray(nums []int) int {
    sum := 0
    maxSum := nums[0]

    for _,num := range nums{
        if sum<0{
            sum = num
        }else{
            sum += num
        }
        
        if sum > maxSum{
            maxSum = sum
        }
    }
    
    return maxSum
}

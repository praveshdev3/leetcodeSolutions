func maxSubArray(nums []int) int {
    max, sum := nums[0], 0
    for _,num := range nums{
        if sum < 0{
            sum = num
        } else{
            sum += num
        }

        if max < sum {
            max = sum
        }
    }
    return max
}

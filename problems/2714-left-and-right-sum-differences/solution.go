func leftRightDifference(nums []int) []int {
    sum := 0
    leftSum := make([]int,len(nums))
    rightSum := make([]int,len(nums))
    result := make([]int,len(nums))

    for index,num := range nums{
        leftSum[index] = sum
        sum += num
    }

    sum = 0
    for i:=len(nums)-1;i>=0;i--{
        rightSum[i] = sum
        sum += nums[i]
    }

    for index,num := range leftSum{
        diff := num - rightSum[index]
        if diff < 0{
            diff = -diff
        }
        result[index] = diff
    }

    return result
}

func sortedSquares(nums []int) []int {
    result := make([]int, len(nums))
    i,j,k := 0,len(nums) - 1,len(nums) - 1

    for i<=j{
        if abs(nums[i]) > abs(nums[j]){
            result[k] = nums[i] * nums[i]
            k--
            i++
        } else{
            result[k] = nums[j] * nums[j]
            k--
            j--
        }
    }

    return result
}

func abs(x int) int{
    if x<=0{
        return -x
    }
    return x
}

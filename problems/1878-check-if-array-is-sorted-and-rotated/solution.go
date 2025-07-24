func check(nums []int) bool {
    if len(nums) <= 1{
        return true
    }

    inversionCount := 0

    for i:=1; i<=len(nums)-1; i++{
        if nums[i] < nums[i-1]{
            inversionCount++
            if inversionCount > 1{
                return false
            }
        }
    }

    if nums[len(nums)-1] > nums[0]{
        inversionCount++
    }

    return inversionCount <=1
}

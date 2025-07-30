func rearrangeArray(nums []int) []int {
    positiveIdx := 0
    negativeIdx := 1
    arr := make([]int,len(nums))
    for i:=0; i<len(nums); i++{
        if nums[i] > 0{
            arr[positiveIdx] = nums[i]
            positiveIdx += 2
        }else{
            arr[negativeIdx] = nums[i]
            negativeIdx += 2
        }
    }
    return arr
}

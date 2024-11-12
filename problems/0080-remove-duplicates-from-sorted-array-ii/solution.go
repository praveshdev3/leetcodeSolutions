func removeDuplicates(nums []int) int {
    index:=1
    occurence:=1

    for i:=1;i<len(nums);i++{
        if nums[i] == nums[i-1]{
            occurence++
        }else{
            occurence=1
        }

        if occurence <= 2 {
            nums[index] = nums[i]
            index++
        }
    }
    return index
}

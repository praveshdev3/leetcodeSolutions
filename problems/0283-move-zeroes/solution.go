func moveZeroes(nums []int)  {
    if len(nums) > 1 {
        j := 0
        for i:=1; i<len(nums); i++{
            if nums[j]==0 && nums[i] !=0{
                nums[i],nums[j] = nums[j],nums[i]
                j++
            }else if nums[j] != 0{
                j++
            }
        }
    }
    // if len(nums) > 1{
    //     j:=0
    //     for i:=1; i<len(nums); i++{
    //         if nums[j] == 0 && nums[i]!=0{
    //             nums[j], nums[i] = nums[i], nums[j]
    //             j++
    //         } else if nums[j] != 0{
    //             j++
    //         }
    //     }
    // }
}

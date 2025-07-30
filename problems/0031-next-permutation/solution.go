func nextPermutation(nums []int) {
    i := len(nums) - 2
    for i>=0 && nums[i+1]<=nums[i]{
        i--
    }

    if i>=0{
        j:=len(nums)-1
        for nums[j] <= nums[i]{
            j--
        }
        nums[j],nums[i]=nums[i],nums[j]
    }

    for l,r:=i+1,len(nums)-1; l<r; l,r=l+1,r-1{
        nums[l],nums[r]=nums[r],nums[l]
    }
}

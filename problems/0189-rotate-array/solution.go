func rotate(nums []int, k int)  {
    k = k%len(nums)
    for i,j:=0,len(nums)-1;i<j;i,j=i+1,j-1{
        nums[i],nums[j] = nums[j],nums[i]
    }
    for i,j:=0,k-1;i<j;i,j=i+1,j-1{
        nums[i],nums[j] = nums[j],nums[i]
    }
    for i,j:=k,len(nums)-1;i<j;i,j=i+1,j-1{
        nums[i],nums[j] = nums[j],nums[i]
    }
}


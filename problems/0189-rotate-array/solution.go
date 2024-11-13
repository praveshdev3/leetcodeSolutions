func rotate(nums []int, k int)  {
    k = k%len(nums)
    for i,j:=0,len(nums)-1;i<j;i,j=i+1,j-1{
        temp := nums[i]
        nums[i] = nums[j]
        nums[j] = temp
    }
    for i,j:=0,k-1;i<j;i,j=i+1,j-1{
        temp := nums[i]
        nums[i] = nums[j]
        nums[j] = temp
    }
    for i,j:=k,len(nums)-1;i<j;i,j=i+1,j-1{
        temp := nums[i]
        nums[i] = nums[j]
        nums[j] = temp
    }
}

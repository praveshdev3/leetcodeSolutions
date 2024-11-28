func productExceptSelf(nums []int) []int {
    ans := make([]int,len(nums))
    ans[0] = 1
    for i:=1;i<=len(nums)-1;i++{
        ans[i] = nums[i-1] * ans[i-1]
    }
    postfix := 1
    for i:=len(nums)-1;i>=0;i--{
        ans[i] = ans[i] * postfix
        postfix = postfix * nums[i]
    }
    return ans
}

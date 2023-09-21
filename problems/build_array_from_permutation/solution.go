func buildArray(nums []int) []int {
  n := len(nums)
  for i,_ := range nums {
    nums[i] += (n * (nums[nums[i]]%n))
  }
  for i := 0;i < n; i++{
    nums[i]/=n
  }
  return nums
}
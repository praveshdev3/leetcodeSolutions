func canJump(nums []int) bool {
    maxIndex := 0
   for i:=0; i<=len(nums)-1; i++{
    if i > maxIndex{
        return false
    }
    if i + nums[i] > maxIndex{
        maxIndex = i+nums[i]
    }
    if maxIndex >= len(nums) {
        return true
    }
   }
   return true
}

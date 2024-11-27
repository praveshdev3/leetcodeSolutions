func jump(nums []int) int {
    currJump, farthestJump, jumps := 0,0,0
    for i:=0; i<len(nums)-1;i++ {
        if i + nums[i] > farthestJump{
            farthestJump = i + nums[i]
        }
        if currJump == i{
            jumps = jumps+1
            currJump = farthestJump
            if currJump >= len(nums) - 1{
                return jumps
            }
        }
    }
    return 0
}

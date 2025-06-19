func twoSum(nums []int, target int) []int {
    hs := make(map[int]int)
    for i:=0; i<len(nums); i++{
        val,ok := hs[target - nums[i]]
        if ok{
            return []int{val,i}
        }
        hs[nums[i]] = i
    }
    return []int{}
}

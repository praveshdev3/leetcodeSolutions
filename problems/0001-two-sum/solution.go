func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
    for i,num := range nums{
        comp := target - num
        val,ok := mp[comp]
        if ok{
            return []int{val,i}
        }
        mp[num] = i
    }
    return []int{}
}

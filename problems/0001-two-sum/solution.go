func twoSum(nums []int, target int) []int {
    hm := make(map[int]int)
    for i,num := range nums{
        complement := target-num
        if v,ok := hm[complement]; ok{
            return []int{v,i}
        }
        hm[num] = i
    }
    return []int{0,0}
}

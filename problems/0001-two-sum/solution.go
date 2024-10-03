func twoSum(nums []int, target int) []int {
hashmap := make(map[int]int)
for i,num := range nums {
    complement := target - num
    _,ok := hashmap[complement]
    if ok {
        return []int{hashmap[complement],i}
    }
    hashmap[num] = i
}
return nil
}

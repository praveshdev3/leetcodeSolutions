func removeDuplicates(nums []int) int {
     mp := make(map[int]bool)
    k := 0

    for _,num := range nums {
        if _,ok := mp[num]; !ok {
            mp[num] = true
            nums[k] = num
            k++
        }
    }

    return k
}

func smallerNumbersThanCurrent(nums []int) []int {
    sortedNums := make([]int,len(nums))
    copy(sortedNums,nums)
    sort.Ints(sortedNums)
    counts := make([]int,len(nums))
    cache := map[int]int{}
    prevN := -1

    for i,n := range sortedNums {
        if n != prevN {
            cache[n] = i
        }
        prevN = n
    }

    for i,n := range nums {
        counts[i],_ = cache[n]
    }
    
    return counts
}

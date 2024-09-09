func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    count := 0
    for _,num := range nums1{
        for _,n := range nums2{
            if num % (n*k) == 0{
                count++
            }
        }
    }
    return count
}

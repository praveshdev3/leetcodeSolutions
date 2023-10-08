func numIdenticalPairs(nums []int) int {
    ans := 0
    cnt := make(map[int]int)
    for _,x := range nums {
        ans += cnt[x]
        cnt[x]++
    }
    return ans
}

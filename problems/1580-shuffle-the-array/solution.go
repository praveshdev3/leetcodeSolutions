func shuffle(nums []int, n int) []int {
    resp := make([]int,2*n)
    for i := 0; i<n; i++ {
        resp[i+i] = nums[i]
        resp[i+i+1] = nums[n+i]
    }
    return resp
}

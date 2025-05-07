func countBits(n int) []int {
    ans := make([]int, n+1)
    offset := 1

    for i := 1; i <= n; i++ {
        if offset*2 == i {
            offset *= 2
        }
        ans[i] = ans[i - offset] + 1
    }

    return ans
}

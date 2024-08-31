func findMatrix(nums []int) [][]int {
    freq := make(map[int]int)
    ans := make([][]int,0)

    for _, value := range nums{
        if freq[value] >= len(ans){
            ans = append(ans,[]int{})
        }

        ans[freq[value]] = append(ans[freq[value]], value)
        freq[value]++
    }

    return ans
}

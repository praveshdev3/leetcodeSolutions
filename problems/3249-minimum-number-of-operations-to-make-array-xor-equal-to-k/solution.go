func minOperations(nums []int, k int) int {
    xorOfNums := 0
    for _,num := range nums{
        xorOfNums = xorOfNums ^ num
    }
    return bits.OnesCount(uint(xorOfNums ^ k))
}

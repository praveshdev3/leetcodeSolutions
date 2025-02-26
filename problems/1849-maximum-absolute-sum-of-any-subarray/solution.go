func maxAbsoluteSum(nums []int) int {
    var minPrefixSum float64 = 0
    var  maxPrefixSum float64 = 0
    var prefixSum float64 = 0

    for _, num := range nums{
        prefixSum += float64(num)
        minPrefixSum = math.Min(minPrefixSum, prefixSum)
        maxPrefixSum = math.Max(maxPrefixSum, prefixSum)
    }

    return int(maxPrefixSum-minPrefixSum)
}

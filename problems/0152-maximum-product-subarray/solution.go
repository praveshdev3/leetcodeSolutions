func maxProduct(nums []int) int {
    maxProduct,currMax,currMin := nums[0],nums[0],nums[0]
    product := 1
    for i:=1; i<len(nums); i++{
        num := nums[i]
        product = max(num, max(num*currMax, num*currMin))
        currMin = min(num, min(num*currMax, num*currMin))
        currMax = product
        maxProduct = max(maxProduct, currMax)
    }
    return maxProduct
}

func max(a,b int) int{
    if a>b{
        return a
    }
    return b
}

func min(a,b int) int{
    if a<b{
        return a
    }
    return b
}

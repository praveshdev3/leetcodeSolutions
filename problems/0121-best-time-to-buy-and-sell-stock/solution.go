func maxProfit(prices []int) int {
    maxProfit := 0
    minPrice := prices[0]
    for i:=1; i<len(prices); i++{
        if prices[i] - minPrice > maxProfit {
            maxProfit = prices[i] - minPrice
        }
        if prices[i]<minPrice{
            minPrice = prices[i]
        }
    }
    return maxProfit
}

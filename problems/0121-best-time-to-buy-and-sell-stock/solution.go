func maxProfit(prices []int) int {
    maxProfit := 0
    buyPrice := prices[0]
    for i:=1; i<len(prices); i++{
        if prices[i] < buyPrice{
            buyPrice = prices[i]
        }
        profit := prices[i] - buyPrice
        if maxProfit < profit {
            maxProfit = profit
        }
    }
    return maxProfit
}

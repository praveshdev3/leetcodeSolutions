func maxProfit(prices []int) int {
    maxProfit := 0
    buyPrice := prices[0]
    for i:=1; i<len(prices);i++{
        if prices[i] < buyPrice{
            buyPrice = prices[i]
        } else{
            profit := prices[i] - buyPrice
            if profit > maxProfit{
                maxProfit = profit
            }
        }
    }
    return maxProfit
}

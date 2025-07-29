func maxProfit(prices []int) int {
    maxProfit := 0
    buyPrice := prices[0]
    for i:=1; i<len(prices); i++{
        profit := 0
        if prices[i] > buyPrice{
            profit = prices[i] - buyPrice
        }else{
            buyPrice = prices[i]
        }
        if profit > maxProfit{
            maxProfit = profit
        }
    }
    return maxProfit
}

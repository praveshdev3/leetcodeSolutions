func maxProfit(prices []int) int {
    profit := 0
    minPrice := prices[0]

    for _,price := range prices{
        if price < minPrice{
            minPrice = price
        }
        if price - minPrice > profit{
            profit = price-minPrice
        }
    }

    return profit
}

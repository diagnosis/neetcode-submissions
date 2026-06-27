func maxProfit(prices []int) int {
    n := len(prices)
    maxProfit := 0
     minPrice := prices[0]
for i := 1 ;i < n ; i ++ {
    if minPrice > prices[i]{
        minPrice = prices[i]
    }else{
        if maxProfit < prices[i] - minPrice{
            maxProfit = prices[i] - minPrice
        }
    }
}

return maxProfit
}

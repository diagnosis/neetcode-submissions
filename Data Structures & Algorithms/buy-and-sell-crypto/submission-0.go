func maxProfit(prices []int) int {
    n := len(prices)
    maxProfit := 0
for i := 0 ;i < n-1 ; i ++ {
    for j := i+1; j < n; j ++ {
        if prices[j] - prices[i] > maxProfit{
            maxProfit = prices[j] - prices[i]
        }
    } 

}
return maxProfit
}

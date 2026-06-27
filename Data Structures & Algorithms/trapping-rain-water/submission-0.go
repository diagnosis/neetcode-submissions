func trap(height []int) int {
    n := len(height)
    maxLeft := make([]int, n)
    maxRight := make([]int, n)
    
    total := 0
    
    for i := 1; i < n; i++{
        maxLeft[i] = max(maxLeft[i-1], height[i-1])
    } 
    for i := n-2; i >=0; i-- {
        maxRight[i] = max(maxRight[i+1], height[i+1])
    }
    for i := 0; i < n; i++ {
        if min(maxLeft[i], maxRight[i]) - height[i] > 0{
            total += min(maxLeft[i], maxRight[i]) - height[i]
        }
    
    }
    return total

}

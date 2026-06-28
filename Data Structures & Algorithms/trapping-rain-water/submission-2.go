func trap(height []int) int {
	water := make([]int, len(height))
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	rightMax[len(height) - 1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := len(height) -2; i >0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i]) 
	}
	for i := 0; i < len(height); i++ {
		w := min(leftMax[i], rightMax[i]) - height[i]
		if w > 0 {
    	water[i] = w
		}
	}
	total := 0 
	for _, v := range water {
		total += v
	}
	return total
	
}

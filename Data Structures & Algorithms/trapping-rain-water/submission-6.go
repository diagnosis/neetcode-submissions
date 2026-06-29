func trap(height []int)int{
	//water := make([]int, len(height))
	leftHeight := make([]int, len(height))
	leftHeight[0] = height[0]
	rightHeight := make([]int, len(height))
	rightHeight[len(height)-1] = height[len(height)-1]
	// left height
	for i := 1; i < len(height); i++ {
		if leftHeight[i-1] < height[i] {
			leftHeight[i] = height[i]
		}else{
			leftHeight[i] = leftHeight[i-1]
		}
	}
	for i := len(height)-2; i >= 0; i-- {
		if rightHeight[i+1] < height[i] {
			rightHeight[i] = height[i]
		}else {
			rightHeight[i] = rightHeight[i+1]
		}
	}
	water := make([]int, len(height))
	water[0] = 0
	water[len(height) -1] = 0
	for i:=1; i < len(height) - 1; i++{
		w := min(leftHeight[i],rightHeight[i]) - height[i]
		if w > 0 {
			water[i] = w
		}
	}
	area := 0
	for _, v := range water{
		area += v
	}
	return area
}
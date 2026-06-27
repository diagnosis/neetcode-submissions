func largestRectangleArea(heights []int) int {
	var stack []int
	maxArea := 0
	for i, h := range heights{
		for len(stack) > 0 && h < heights[stack[len(stack) - 1]]{
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			width :=i
			if len(stack) > 0{
				width = i - stack[len(stack)-1]-1
			}
			area := heights[idx] * width
			maxArea = max(maxArea, area)
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		idx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		width := len(heights)
		if len(stack) > 0 {
			width = len(heights) - stack[len(stack)-1] - 1
		}
		area := heights[idx] * width
		maxArea = max(maxArea, area)
	}

	return maxArea

}
func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	maxArea := (right - left) * min(heights[left], heights[right])
	for left < right {

		if heights[left] < heights[right]{
			left++
			maxArea = max(maxArea, (right - left) * min(heights[left], heights[right]))
		}else{
			right--
			maxArea = max(maxArea, (right - left) * min(heights[left], heights[right]))
		}
	}
	return maxArea

}

func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	area := 0
	for left < right {
		h := min(heights[left], heights[right])
		l := right - left
		area = max(area, l*h)
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return area
}


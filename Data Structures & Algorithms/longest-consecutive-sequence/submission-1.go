func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	var start int
	var count int
	var max int
	for _, num := range nums{
		set[num] = true
	}

	for _, num := range nums{
		if !set[num-1]{
			count = 0
			start = num
			for set[start+count+1] {
			count++
		}
		if count + 1 > max {
			max = count + 1
		}
		}
		
		
	}
	return max
}

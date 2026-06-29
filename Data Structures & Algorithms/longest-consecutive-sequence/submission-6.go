func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums{
		set[num]= struct{}{}
	}

	maxV := 0

	for i := 0 ; i < len(nums); i++ {
		
		if _, ok := set[nums[i]-1]; !ok {
			count := 1
			start := nums[i]
			for{
				if _, ok := set[start+1]; !ok{
					break
				}
				start++
				count++
			}
			maxV = max(maxV, count)
		}

	}
	return maxV

}

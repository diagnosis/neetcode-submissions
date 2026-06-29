func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++{
		set[nums[i]] = struct{}{}
	}
	maxC := 0
	for i := 0 ; i < len(nums); i++ {
		
		if _, ok := set[nums[i]-1]; !ok{
			count := 1
			start := nums[i]
			for {
				if _, ok := set[start+1]; !ok{
					break;
				}
				start++
				count++
			}
			maxC = max(maxC, count)
		}
		
	}
	return maxC
}

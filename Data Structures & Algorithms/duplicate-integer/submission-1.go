func hasDuplicate(nums []int) bool{
	seen := make(map[int]int, len(nums))
	for _,num := range nums {
		seen[num]++
	}
	for _, val := range seen {
		if val > 1 {
			return true
		}
	}
	return false
}

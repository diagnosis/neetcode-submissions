func hasDuplicate(nums []int) bool{
	seen := make(map[int]int, len(nums))
	for _,num := range nums {
		seen[num] = seen[num] + 1
	}
	for _, val := range seen {
		if val > 1 {
			return true
		}
	}
	return false
}

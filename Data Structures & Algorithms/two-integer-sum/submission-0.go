func twoSum(nums []int, target int) []int {
    sumMap := make(map[int]int, len(nums))
	for i, num := range nums {
		comp := target - num
		if idx, ok := sumMap[comp]; ok {
			return []int{idx, i}
		}
		
		sumMap[num] = i
	}
	return []int{}
}

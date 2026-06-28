

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 2)
	for i := 0; i < len(nums); i++{
		c := target - nums[i]
		if _, ok := m[c]; ok{
			res = []int{m[c],i}
			break;
		} 
		m[nums[i]] = i
	}
	return res

}

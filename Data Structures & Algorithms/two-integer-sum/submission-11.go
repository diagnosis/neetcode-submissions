

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 2)
	for i :=0; i < len(nums); i++{
		m[nums[i]] = i
	}
	for i:=0; i < len(nums); i++{
		c := target - nums[i]
		if m[c] == i {
			continue
		}
		if _, ok := m[c]; ok{
			if m[c] < i {
				res = []int{m[c],i}
			}else{
				res = []int{i, m[c]}
			}
			
		}
	}
	return res

}



func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	pair := make([]int,2)
	for i, num := range nums {
		c := target - num
		
		if _ , ok := m[c]; ok{
			pair = []int{m[c],i}
		}
		m[num] = i
		
	}
	return pair

}

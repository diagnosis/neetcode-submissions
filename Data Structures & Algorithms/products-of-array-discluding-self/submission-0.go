func productExceptSelf(nums []int) []int {
		var res []int
	product := func(right, left []int) int {
		s := 1
		t := 1
		for _, n := range right {
			s = s * n
		}
		for _, n := range left {
			t = t * n
		}
		return s * t

	}

	for i, _ := range nums {
		res = append(res, product(nums[:i], nums[i+1:]))
	}
	return res
}

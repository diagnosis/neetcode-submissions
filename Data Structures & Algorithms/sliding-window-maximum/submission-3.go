func maxSlidingWindow(nums []int, k int) []int {
	dq := make([]int, 0)
	start := 0
	res := make([]int, 0)
	for end:= 0 ;end < len(nums);end ++ {
		
		for len(dq) > 0 && nums[end] > nums[dq[len(dq)-1]]{
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, end)

		start = end-k+1
		if start > dq[0]{
			dq = dq[1:]
		}
		if start >= 0 {
			res = append(res, nums[dq[0]])
		}


	} 
	return res


    
}

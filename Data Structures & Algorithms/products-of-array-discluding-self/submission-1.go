func productExceptSelf(nums []int) []int {
	res := make([]int,0)
	
	
	for i := 0; i<len(nums); i++ {
		leftProduct := 1
	rightProduct := 1
		for j := i-1; j >= 0; j-- {
			leftProduct = leftProduct *  nums[j]
		}
		for k := i+1; k < len(nums); k++{
			rightProduct = rightProduct * nums[k]
		}
		res = append(res, leftProduct*rightProduct)
	}
	return res

}

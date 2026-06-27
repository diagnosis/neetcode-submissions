func productExceptSelf(nums []int) []int {
	res := make([]int,len(nums))
	leftProduct := 1
	for i :=0; i<len(nums); i++{
		res[i] = leftProduct
		leftProduct *= nums[i]
	}
	rightProduct := 1
	for i := len(nums) -1; i >=0; i--{
		res[i] = rightProduct * res[i]
		rightProduct *= nums[i]
	}
	
	return res

}

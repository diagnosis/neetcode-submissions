func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))
	leftProducts := make([]int, len(nums))
	rightProducts := make([]int, len(nums))
	leftProducts[0] = 1
	rightProducts[len(nums)-1] = 1
	for i := 1; i < len(nums); i++{
		leftProducts[i] = nums[i-1] * leftProducts[i-1]
	}
	for i := len(nums) - 2; i >= 0; i--{
		rightProducts[i] = nums[i+1] * rightProducts[i+1]
	}
	for i := range nums {
		products[i] = leftProducts[i] * rightProducts[i]
	}
	return products
}

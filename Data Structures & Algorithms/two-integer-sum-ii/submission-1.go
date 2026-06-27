func twoSum(nums []int, target int) []int {
	left := 0 
	right := len(nums) - 1
	res := make([]int, 0, 2)
	for left < right {
		if nums[left] + nums[right] == target{
			res = append(res, []int{left+1, right+1}...)
			break;
		}
		if nums[left] + nums[right] > target{
			right--
		}
		if nums[left] + nums[right] < target{
			left++
		}
	}
	return res
}

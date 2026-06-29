func twoSum(nums []int, target int) []int {
	
	left := 0
	right := len(nums) - 1
	res := make([]int,2)
	for left < right {
		if nums[left] + nums[right] < target{
			left++
		}else if nums[left] + nums[right] > target{
			right--
		}else{
			res[0] = left + 1
			res[1] = right + 1
			break;
		}
	}
	return res


}

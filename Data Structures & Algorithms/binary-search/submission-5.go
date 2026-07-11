func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		med := (left + right) / 2
		if nums[med] == target {
			return med
		} else if nums[med] > target {
			right = med - 1
		} else {
			left = med + 1
		}
	}
	return -1
}

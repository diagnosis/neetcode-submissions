func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	mid := (end + start) / 2
	for start <= end {
		if nums[mid] > target{
			end = mid - 1
		}else if nums[mid] < target{
			start = mid + 1
		}else {
			return mid
		}
		mid = (end + start) / 2
	}
	return -1
}

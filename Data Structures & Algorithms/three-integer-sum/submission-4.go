import "slices"
func threeSum(nums []int) [][]int {
	var res [][]int
	slices.Sort(nums)
	
	for i := 0; i < len(nums); i ++ {
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1
		if i>0 && nums[i] == nums[i-1]{
			continue
		}
		for left < right {
			if nums[left] + nums[right] > target{
				right --
			}else if nums[left] + nums[right] < target {
				left ++
			}else {
			
				res = append(res, []int{-target, nums[left], nums[right]})
				right -- 
				left ++
				for left < right && nums[left] == nums[left-1] {
    left++
}
for left < right && nums[right] == nums[right+1] {
    right--
}
			
			}
		}
		
	}
	return res
}

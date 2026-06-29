
import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int,0)
	
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
    		continue
		}

		target := -nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right{
			if nums[left] + nums[right] < target{
				left++
			}else if nums[left] + nums[right] > target{
				right--
			}else{
				res = append(res, []int{nums[left], nums[right], -target})
				left++
				right--
				for left < right && nums[left] == nums[left - 1]{left++}
				for left < right && nums[right] == nums[right+1]{right--}
			}

		}
	}
	return res

}

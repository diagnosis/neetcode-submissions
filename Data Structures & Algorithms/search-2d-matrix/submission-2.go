func searchMatrix(matrix [][]int, target int) bool {
	var search1 = func(nums []int, target int) bool {
		left := 0
		right := len(nums) - 1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				return true
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return false
	}

	top, bottom := 0, len(matrix)-1
	for top <= bottom {
		mid := (top + bottom) / 2
		row := matrix[mid]

		if target < row[0] {
			bottom = mid - 1
		} else if target > row[len(row)-1] {
			top = mid + 1
		} else {
			return search1(row, target)
		}
	}
	return false

}

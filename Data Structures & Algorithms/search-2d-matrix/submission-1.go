func searchMatrix( matrix [][]int, target int)bool{
	for _, row := range matrix{
		start := 0
		end := len(row) - 1
		mid := (end + start) / 2
		for start <= end{
			if target < row[mid]{
				end = mid -1
			}else if target > row[mid]{
				start = mid + 1
			}else{
				return true
			}
			mid = (end + start) / 2
		}
	}
	return false
}


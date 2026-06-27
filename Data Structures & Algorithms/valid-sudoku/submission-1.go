func isValidSudoku(board [][]byte) bool {
	
	
	// check each row if valid
	for _, row := range board{
		validationMap := make(map[byte]bool)
		for _, v := range row{
			if v == '.'{
				continue
			}
			if validationMap[v]{
				return false
			}
			validationMap[v] = true
		}
	}

	//check each column
	for i, _ := range board{
		validationMap := make(map[byte]bool)
		for j := 0; j < len(board); j++ {
			if board[j][i] == '.'{
				continue
			}
			if validationMap[board[j][i]]{
				return false
			}
			validationMap[board[j][i]] = true
		}
	}


	//check each subboxes
	for boxRow := 0; boxRow < 3; boxRow++{
		for boxCol := 0; boxCol < 3; boxCol++{
			validationMap := make(map[byte]bool)
			for r := 0; r < 3; r ++ {
				for c := 0; c < 3; c ++ {
					actualRow := boxRow * 3 + r
					actualCol := boxCol * 3 + c
					v := board[actualRow][actualCol]
					if v == '.'{
						continue
					}
					if validationMap[v]{
						return false
					}
					validationMap[v] = true
				}
			}
		}
	}
	return true


}

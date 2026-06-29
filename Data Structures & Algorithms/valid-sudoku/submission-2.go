func isValidSudoku(board [][]byte) bool {
	
	//check row 
	for i := 0; i < 9; i ++ {
		rowMap := make(map[byte]bool)
		for j := 0; j < 9; j++{
			if board[i][j] == '.'{
				continue
			}
			if _, ok := rowMap[board[i][j]]; ok{
				return false
			}
			rowMap[board[i][j]] = true
		}
	}
	//check column
	for i := 0; i < 9; i ++ {
		colMap := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.'{
				continue
			}
			if _, ok := colMap[board[j][i]]; ok {
				return false
			}
			colMap[board[j][i]] = true
		}
	}

	// check squares
	for oi := 0; oi < 3; oi ++ {
		for oj := 0; oj < 3; oj++ {
			subMap := make(map[byte]bool)
		
			for i:=3*(oi); i < 3+3*oi ; i++{
				for j := 3*(oj); j < 3+3*oj; j++{
					if board[i][j] == '.'{
						continue
					}
					if _, ok := subMap[board[i][j]]; ok {
						return false
					}
					subMap[board[i][j]] = true
				}
			}
		}
	}

	return true


}

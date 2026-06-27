func isValidSudoku(board [][]byte) bool {
    //check row and cols
    l := len(board)
    for i := 0; i < l ; i ++ {
        colMap := make(map[byte]bool)
        rowMap := make(map[byte]bool)
        for j := 0; j < l; j ++ {
            valRow := board[i][j]
            valCol := board[j][i]
            if valRow != '.'{
                if _, ok := rowMap[valRow]; ok {
                    return false
                } 
                rowMap[valRow] = true
            }
            if valCol != '.'{
                if _, ok := colMap[valCol]; ok {
                    return false
                }
                colMap[valCol] = true
            }
        }
    }
    // check subBox
    for i :=0; i < l; i += 3{
        for j := 0; j < l; j +=3 {
            subMap := make(map[byte]bool)
            for innerI := i; innerI < i+3; innerI++{
                for innerJ := j; innerJ < j+3; innerJ++{
                    val := board[innerI][innerJ]
                    if val != '.'{
                        if _, ok := subMap[val]; ok {
                            return false
                        }
                        subMap[val]= true
                    }
                }
            }
        }
    }

    return true


}

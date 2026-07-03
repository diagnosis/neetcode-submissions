func isValid(s string) bool {
    pMap := make(map[rune]rune)
	pMap[')'] = '('
	pMap['}'] = '{'
	pMap[']'] = '['
	
	stack := make([]rune, 0)

	for _, r := range s {
		if _, ok := pMap[r]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != pMap[r]{
				return false
			}
			stack = stack[:len(stack)-1]
		}else{
			stack = append(stack, r)
		}
	}

	return len(stack) == 0


}

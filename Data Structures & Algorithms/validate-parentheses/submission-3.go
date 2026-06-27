

func isValid(s string) bool {
	pMap := map[rune]rune{
		'[':']',
		'{':'}',
		'(':')',
	}
	stack := make([]rune, 0)
	for _, r := range s {
		if r == '(' || r == '{' || r == '['{
			stack = append(stack, r)
		}else {
			if len(stack) == 0 || pMap[stack[len(stack)-1]] != r  {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}



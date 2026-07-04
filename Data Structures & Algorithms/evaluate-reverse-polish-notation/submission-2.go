func evalRPN(tokens []string) int {
	opMap := map[string]struct{}{
		"+": {},
		"-": {},
		"/": {},
		"*": {},
	}
	stack := make([]int, 0)
	operation := ""
	res := 0
	for i := 0; i < len(tokens); i++ {
		num, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, num)
		} else {
			operation = tokens[i]
		}
		if _, ok := opMap[tokens[i]]; ok {
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch operation {
			case "+":
				res = num2 + num1
			case "-":
				res = num2 - num1
			case "*":
				res = num2 * num1
			case "/":
				res = num2 / num1
			}
			stack = append(stack, res)
		}else{
            res = stack[len(stack)-1]
        }
	}
	return res

}
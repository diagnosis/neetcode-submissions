func evalRPN(tokens []string) int {
    stack := make([]int, 0)
    res := 0
    for _, s := range tokens {
       n, err := strconv.Atoi(s)
       if err == nil {
        stack = append(stack, n)
       }else{
        switch s {
            case "+": {
                res = stack[len(stack)-1]  + stack[len(stack)-2] 
                stack = stack[:len(stack)-2]
                stack = append(stack, res)
            }
            case "-": {
                res = stack[len(stack)-2]  - stack[len(stack)-1] 
                stack = stack[:len(stack)-2]
                stack = append(stack, res)
            }
            case "*": {
                res = stack[len(stack)-1]  * stack[len(stack)-2] 
                stack = stack[:len(stack)-2]
                stack = append(stack, res)
            }
            case "/": {
                res = stack[len(stack)-2]  / stack[len(stack)-1] 
                stack = stack[:len(stack)-2]
                stack = append(stack, res)
            }
        }
       }
    }
    return stack[0]
}

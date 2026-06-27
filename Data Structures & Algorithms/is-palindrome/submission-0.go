func isPalindrome(s string) bool {
    l := strings.ToLower(s)
    result := make([]byte, 0, len(s))
    for i :=0; i < len(s); i++{
        c := l[i] 
        if (c >= 'a' && c <='z') || (c >= '0' && c<= '9'){
            result = append(result, c)
        }
    }
   
    reversed := make([]byte, 0, len(s))
    for i := len(result)-1; i >=0; i-- {
        reversed = append(reversed, result[i])
    }
    
    if string(reversed) == string(result){
        return true
    }
    return false
}

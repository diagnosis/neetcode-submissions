func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if !(unicode.IsDigit(rune(s[left])) || unicode.IsLetter(rune(s[left]))){
			left++
			continue
		}
		if !(unicode.IsDigit(rune(s[right])) || unicode.IsLetter(rune(s[right]))){
			right--
			continue
		}
		if unicode.ToLower(rune(s[right])) != unicode.ToLower(rune(s[left])){
			return false
		}
		left++
		right--
	}
	return true
}

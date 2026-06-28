func lengthOfLongestSubstring(s string) int {
	left := 0 
	seen := make(map[byte]bool)
	maxLen := 0
	for right := 0; right < len(s); right++ {
		for seen[s[right]]{
			delete(seen, s[left])
			left++
		}
		seen[s[right]]= true
		maxLen = max(maxLen, right - left + 1)
	}

	return maxLen
}

func lengthOfLongestSubstring(s string) int {
	start := 0
	window := 0
	seen := make(map[byte]struct{})

	for end:=0; end < len(s); end++{
		for {
			if _, ok := seen[s[end]]; !ok {
				break;
			}
			delete(seen, s[start])
			start++
		}
		seen[s[end]]= struct{}{}
		window = max(window, end-start+1)
	}
	return window
}

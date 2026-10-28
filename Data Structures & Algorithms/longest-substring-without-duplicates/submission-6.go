func lengthOfLongestSubstring(s string) int {
	start := 0
	window := 0
	seen := make(map[rune]struct{})
	for i, c := range s {
		for {
			if _, ok := seen[c]; !ok {
				break
			}
			delete(seen, rune(s[start]))
			start++
		}
		seen[c] = struct{}{}
		l := i + 1 - start
		window = max(window, l)
	}
	return window
}

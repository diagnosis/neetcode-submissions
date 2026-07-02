func minWindow(s string, t string) string {
	mapT := make(map[byte]int)
	window := make(map[byte]int)
	minLen := len(s)+1
	start := 0 
	have := 0
	res := ""
	for i := 0; i < len(t); i++ {
		mapT[t[i]]++
	}
	need := len(mapT)

	for end := 0; end < len(s); end++ {
		window[s[end]]++
		if window[s[end]] == mapT[s[end]]{
			have++
		}

		for have == need{
			currentLen := end - start + 1
			if currentLen < minLen{
				minLen = currentLen
				res = s[start:start+minLen]
			}
			leaving := s[start]
			window[leaving]-- 
			if mapT[leaving] > window[leaving]{
				have--
			}

			if window[leaving] == 0 {
				delete(window, leaving)
			}
			start++

		}
	}

	return res


}

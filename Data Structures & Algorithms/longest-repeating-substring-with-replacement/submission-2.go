func characterReplacement(s string, k int) int {
	left := 0
	windowSize := 0
	freqMap := make(map[byte]int)
	maxFreq := 0
	for right := 0; right < len(s); right++ {
		freqMap[s[right]]++
		maxFreq = max(maxFreq, freqMap[s[right]])	
		if (right - left + 1) - maxFreq > k {
			freqMap[s[left]]--
			left++
		}
		windowSize = max(windowSize, right-left + 1)
		
	}
	return windowSize

}

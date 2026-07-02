func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	maxFreq := 0 
	start := 0 
	maxLen := 0 

	for end := 0; end < len(s); end++ {
		freqMap[s[end]]++ 
		maxFreq = max(maxFreq, freqMap[s[end]])
		if end - start + 1 > k + maxFreq{
			freqMap[s[start]]--
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}

	return maxLen

}

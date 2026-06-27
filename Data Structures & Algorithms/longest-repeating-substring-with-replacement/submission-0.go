func characterReplacement(s string, k int) int {
    var left, right, maxCount, maxLen int
    chMap := make(map[rune]int)
    for i, c := range s {
        right = i 
        chMap[c]++
        maxCount = max(maxCount, chMap[c])

        if (right -left +1) - maxCount > k{
            chMap[rune(s[left])]--
            left++
        }
        maxLen = max(maxLen, right- left +1)


    }
    return maxLen
    
}

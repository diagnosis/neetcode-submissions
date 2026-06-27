func minWindow(s string, t string) string {
    tMap := make(map[rune]int)
    for _ ,c := range t {
        tMap[c]++
    }
    need := len(tMap)
    have := 0
    left := 0
    windowMap := make(map[rune]int)
    minLen := math.MaxInt32
    result := ""
    for i, c := range s {
        windowMap[c]++
        if windowMap[c] == tMap[c]{
            have++
        }
        for have == need {
            
            if i - left +1 < minLen{
                 minLen = i - left + 1
                result = s[left:i+1]
            }
            windowMap[rune(s[left])]--
            if windowMap[rune(s[left])] < tMap[rune(s[left])]{
                have--
            }
            left ++
           
        }
    }
    return result
    
}

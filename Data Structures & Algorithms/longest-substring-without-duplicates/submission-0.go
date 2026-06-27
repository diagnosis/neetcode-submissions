func lengthOfLongestSubstring(s string) int {
   right := 0
   left := 0
   window :=0
   chMap := make(map[rune]int)
   for i, ch := range s {
    right = i
    if _, ok := chMap[ch]; ok {
        left = max(left, chMap[ch] + 1) 
    }
    chMap[ch] = i
    if window < right - left + 1{
            window = right - left + 1
    }
    
   }
   return window

}

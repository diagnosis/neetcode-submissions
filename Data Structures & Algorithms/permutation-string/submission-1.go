import "maps"
func checkInclusion(s1 string, s2 string) bool {
    s1Len := len(s1)
    s1Map := make(map[rune]int)
    s2Map := make(map[rune]int)
    for _, c := range s1 {
        s1Map[c]++
    }
    for i, c := range s2{
        s2Map[c]++

        if i >= s1Len{
            s2Map[rune(s2[i-s1Len])]--
            if s2Map[rune(s2[i-s1Len])] == 0 {
                delete(s2Map, rune(s2[i-s1Len])) 
            }
        }
          if maps.Equal(s2Map, s1Map){
                return true
            }
    }
    return false
    

}

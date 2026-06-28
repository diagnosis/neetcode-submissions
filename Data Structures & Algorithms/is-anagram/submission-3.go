
func isAnagram(s string, t string) bool {
	mapS := make(map[byte]int)
	mapT := make(map[byte]int)
	for i :=0; i < len(s); i++ {
		mapS[s[i]]++
	}
	for j := 0; j < len(t); j++ {
		mapT[t[j]]++
	}
	if len(mapS) != len(mapT){
		return false
	}
	for i, v := range mapS {
		if v != mapT[i]{
			return false
		}
	}
	return true
}

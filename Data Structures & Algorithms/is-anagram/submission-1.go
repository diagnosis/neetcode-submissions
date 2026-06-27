import "maps"
func isAnagram(s string, t string) bool {
	ms := make(map[byte]int)
	mt := make(map[byte]int)
	if len(s) != len(t){
		return false
	}
	for i := 0 ; i < len(s); i ++ {
		ms[s[i]]++
		mt[t[i]]++
	}
	if maps.Equal(ms,mt){
		return true
	}
	return false

}

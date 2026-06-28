
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	mapI := make(map[byte]int)
	for i :=0; i < len(s); i++ {
		mapI[s[i]]++
	}
	for i := 0; i < len(t); i ++ {
		if _, ok := mapI[t[i]]; !ok {
			return false
		}
		mapI[t[i]]--
		if mapI[t[i]] == 0{
			delete(mapI, t[i])
		}
	}
	return true
}

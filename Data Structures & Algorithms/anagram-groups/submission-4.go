import "slices"
func groupAnagrams(strs []string) [][]string {
	
	resMap := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		resMap[string(b)] = append(resMap[string(b)], str)
	}
	resSlice := make([][]string,0)
	for _, v := range resMap{
		resSlice = append(resSlice, v)
	}
	return resSlice

}

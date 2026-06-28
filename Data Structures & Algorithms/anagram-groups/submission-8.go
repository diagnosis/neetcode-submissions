
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for i := 0; i < len(strs); i++ {
		loc := [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			loc[int(strs[i][j]-97)]++
		}
		m[loc] = append(m[loc], strs[i])
	}
	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res

}

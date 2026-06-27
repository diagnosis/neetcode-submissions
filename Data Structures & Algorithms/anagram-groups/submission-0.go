func groupAnagrams(strs []string) [][]string {
	sortString := func(s string) string {
		runeSlice := []rune(s)
		sort.Slice(runeSlice, func(i, j int) bool {
			return runeSlice[i] < runeSlice[j]
		})
		return string(runeSlice)
	}

	m := make(map[string][]string)
	for _, s := range strs {
		key := sortString(s)
		m[key] = append(m[key], s)
	}
	
	res := make([][]string, 0, len(m))
	
	for _, g := range m {
		res = append(res, g)
	} 
	return res
}

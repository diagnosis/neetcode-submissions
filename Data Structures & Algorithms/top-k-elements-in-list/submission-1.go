func topKFrequent(nums []int, k int) []int {
	
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys[:k]


}

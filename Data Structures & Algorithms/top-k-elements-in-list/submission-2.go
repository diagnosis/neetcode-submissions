func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, n := range nums {
		freqMap[n]++
	}
	bucket := make([][]int, len(nums)+1)
	for key, v := range freqMap{
		bucket[v] = append(bucket[v], key)
	}
	res := make([]int,0,k)
	for i:= len(bucket)-1; i>=0; i--{
		for _, v := range bucket[i]{
			if len(res) == k{
				break;
			}
			res = append(res, v)
		}
	}
	return res

}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _ , num := range nums {
		freqMap[num]++
	}

	bucket := make([][]int,len(nums)+1)

	for key, v := range freqMap {
		bucket[v]= append(bucket[v], key)
	}

	res := make([]int,0, k)
	for i := len(bucket) - 1; i >= 0; i --{
		if len(res) == k {
			break;
		}
		for j := 0; j < len(bucket[i]); j++{
			if len(res) == k {
				break;
			}
			res = append(res, bucket[i][j])
		}
		
	}
	return res
}

func topKFrequent(nums []int, k int) []int {
	//first define freq map, and bucket

	freqMap := make(map[int]int)
	bucket := make([][]int, len(nums)+1)
	
	//loop to fill map
	for _, v := range nums {
		freqMap[v]++
	}

	//bucket sort

	for key, v := range freqMap {
		bucket[v] = append(bucket[v], key)
	}

	//now generate results 
	res := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0 ; i--{
		if len(res) == k {
			break
		}
		for _, v := range bucket[i]{
			if len(res) == k {
				break
			}
			res = append(res, v)
		}
	}
	return res

}

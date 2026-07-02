
import "maps"

func checkInclusion(s1 string, s2 string) bool {
  	mapS1 := make(map[byte]int)
	for i :=0; i < len(s1); i ++ {
		mapS1[s1[i]]++
	}
	window := make(map[byte]int)
	start := 0 


	for end :=0 ; end <  len(s2); end++ {
		window[s2[end]]++
		for len(s1) == end - start + 1{
			if maps.Equal(window, mapS1){
				return true
			}
			
				window[s2[start]]--
				if window[s2[start]] == 0{
					delete(window, s2[start])
				}
			
			start++
		}
	}
	return false
}

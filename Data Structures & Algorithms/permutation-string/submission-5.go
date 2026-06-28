
import "maps"

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2){
		return false
	}
	mapS1 := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(s1); i++{
		mapS1[s1[i]]++
		window[s2[i]]++
	}

	if maps.Equal(mapS1, window){
		return true
	}

	for right := len(s1); right < len(s2); right++{
		window[s2[right]]++
		leaving := s2[right - len(s1)]
		window[leaving]--
		if window[leaving] == 0{
			delete(window, leaving)
		}
		if maps.Equal(mapS1, window){
			return true
		}

	}
	return false

	

}

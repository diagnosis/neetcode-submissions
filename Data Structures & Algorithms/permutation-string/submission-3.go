
import "maps"

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    mapS1 := make(map[rune]int)
    for _, b := range s1 {
        mapS1[b]++
    }
    for i := 0; i+len(s1) <= len(s2); i++ {
        mapWindow := make(map[rune]int)
        for _, b := range s2[i : i+len(s1)] {
            mapWindow[b]++
        }
        if maps.Equal(mapS1, mapWindow) {
            return true
        }
    }
    return false
}

import (
	"slices"
)
func minEatingSpeed(piles []int, h int) int{
	mid := 0
	end := slices.Max(piles)
	start := 1
	calcPilesHours := func(piles []int, rate int)int{
		sum := 0
		for _, pile := range piles {
			sum = sum + int(math.Ceil(float64(pile)/float64(rate)))
		}
		return sum
	}

	for start < end {
		mid = (start + end) / 2
		if calcPilesHours(piles, mid) > h{
			start = mid + 1
		}else{
			end = mid
		}
	}
	return start
}

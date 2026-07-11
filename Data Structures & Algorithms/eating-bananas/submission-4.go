import (
	"slices"
)

func minEatingSpeed(piles []int, h int) int {

	maxRate := slices.Max(piles)
	minRate := 1
	for minRate <= maxRate {
		midRate := (minRate + maxRate) / 2
		totalH := 0
		for _, pile := range piles {
			totalH += int(math.Ceil(float64(pile) / float64(midRate)))
		}
		if totalH > h {
			minRate = midRate + 1
		} else {
			maxRate = midRate - 1
		}

	}
	return minRate

}
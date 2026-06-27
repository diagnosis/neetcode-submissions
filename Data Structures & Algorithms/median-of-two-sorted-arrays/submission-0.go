import "slices"
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := slices.Concat(nums1, nums2)
	slices.Sort(nums)
	l := len(nums)
	if l % 2 == 0 {
		return (float64(nums[l/2]) + float64(nums[l/2 -1]))/2
	}

	return float64(nums[(l-1)/2])
}

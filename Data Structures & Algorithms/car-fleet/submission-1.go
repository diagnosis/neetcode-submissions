func carFleet(target int, position []int, speed []int) int {
	indices := make([]int, len(position))
	for i := range position {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return position[indices[i]] > position[indices[j]]
	})

	times := make([]float64, len(position))

	for i, vI := range indices {
		times[i] = float64(target-position[vI]) / float64(speed[vI])
	}
	myStack := make([]float64, 0)
	for i := 0; i < len(times); i++ {
		if len(myStack) == 0 || times[i] > myStack[len(myStack)-1] {
			myStack = append(myStack, times[i])
		}
	}
	return len(myStack)

}

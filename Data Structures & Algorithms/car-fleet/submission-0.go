func carFleet(target int, positions []int, speeds []int) int {
	timeCalc := func (position, speed int)float64{
		return float64(target-position) / float64(speed)
	}
	indices := make([]int, len(positions))
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int)bool{
		return positions[indices[i]] > positions[indices[j]]
	})
	stack := make([]float64, 0)
	for _, idx := range indices {
		time := timeCalc(positions[idx], speeds[idx])
		// now stack logic here
		if len(stack) == 0 || time > stack[len(stack)-1]{
			stack = append(stack, time)
		}
	}

	return len(stack)
}

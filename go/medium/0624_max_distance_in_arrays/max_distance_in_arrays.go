package maxdistanceinarrays

func maxDistance(arrays [][]int) int {
	minVal := arrays[0][0]
	maxVal := arrays[0][len(arrays[0])-1]

	distance := 0
	for i := 1; i < len(arrays); i++ {
		n := len(arrays[i])
		distance = max(
			distance,
			intAbs(arrays[i][n-1]-minVal),
			intAbs(arrays[i][0]-maxVal),
		)

		minVal = min(minVal, arrays[i][0])
		maxVal = max(maxVal, arrays[i][n-1])
	}

	return distance
}

func intAbs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

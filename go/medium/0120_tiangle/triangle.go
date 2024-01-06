package tiangle

import "slices"

func minimumTotal(triangle [][]int) int {
	levels := make([][]int, 0, len(triangle))
	for lvl, row := range triangle {
		levels = append(levels, make([]int, len(row)))
		for i := range levels[lvl] {
			levels[lvl][i] = 1 << 32
		}
	}
	levels[0][0] = triangle[0][0]

	for i := 0; i < len(triangle)-1; i++ {
		row, nextRow := levels[i], triangle[i+1]
		for j := 0; j < len(row); j++ {
			levels[i+1][j] = min(levels[i+1][j], row[j]+nextRow[j])
			levels[i+1][j+1] = min(levels[i+1][j+1], row[j]+nextRow[j+1])
		}
	}

	return slices.Min(levels[len(levels)-1])
}

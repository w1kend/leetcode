package minimumpathsum

func minPathSum(grid [][]int) int {
	for i, row := range grid {
		for j := range row {
			switch {
			case j > 0 && i > 0:
				// can move from both sides
				grid[i][j] = min(grid[i][j]+grid[i][j-1], grid[i][j]+grid[i-1][j])
			case j > 0:
				// only horizontally
				grid[i][j] = grid[i][j] + grid[i][j-1]
			case i > 0:
				// only vertically
				grid[i][j] = grid[i][j] + grid[i-1][j]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

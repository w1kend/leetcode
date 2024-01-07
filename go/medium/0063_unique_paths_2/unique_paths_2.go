package uniquepaths2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	pathsGrid := make([][]int, 0, len(obstacleGrid))
	for _, row := range obstacleGrid {
		pathsGrid = append(pathsGrid, make([]int, len(row)))
	}
	pathsGrid[0][0] = 1

	for i := range pathsGrid {
		for j := range pathsGrid[i] {
			if obstacleGrid[i][j] == 1 {
				// do nothing. 0 ways to go here
				continue
			}
			pathsGrid[i][j] =
				pathsTo(pathsGrid, obstacleGrid, i-1, j) + pathsTo(pathsGrid, obstacleGrid, i, j-1)
		}
	}

	return pathsGrid[len(pathsGrid)-1][len(pathsGrid[0])-1]
}

func pathsTo(pathsGrid, obstacleGrid [][]int, i, j int) int {
	if i == 0 && j == -1 {
		// special case for the first cell
		return 1
	}

	if i < 0 || i >= len(pathsGrid) || j < 0 || j >= len(pathsGrid[i]) {
		return 0
	}

	if obstacleGrid[i][j] == 1 {
		return 0
	}

	return pathsGrid[i][j]
}

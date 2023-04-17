package numberofislands

func numIslands(grid [][]byte) int {
	islandsNum := 0

	gr := gr(grid)
	for i, row := range gr {
		for j, val := range row {
			if val == '1' {
				islandsNum++
				gr.fillIsland(i, j)
			}
		}
	}

	return islandsNum
}

type gr [][]byte

func (g gr) fillIsland(i, j int) {
	if (i < 0 || i >= len(g)) || (j < 0 || j >= len(g[0])) || g[i][j] == '0' {
		return
	}
	g[i][j] = '0'

	g.fillIsland(i-1, j) // go top
	g.fillIsland(i, j+1) // go right
	g.fillIsland(i+1, j) // go down
	g.fillIsland(i, j-1) // go left
}

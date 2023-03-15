package flood_fill

// https://leetcode.com/problems/flood-fill
// https://leetcode.com/submissions/detail/743656688

//cell, row
type point struct {
	c, r int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 || image[sr][sc] == color {
		return image
	}

	sourceColor := image[sr][sc]

	points := make([]point, 0, len(image))
	points = append(points, point{sc, sr})
	for len(points) > 0 {
		nextPoints := make([]point, 0, 20)
		for _, p := range points {

			image[p.r][p.c] = color

			for _, pp := range neighbours(p, len(image[0]), len(image)) {
				if image[pp.r][pp.c] == sourceColor {
					nextPoints = append(nextPoints, pp)
				}
			}
		}

		points = nextPoints
	}

	return image
}

func neighbours(p point, xlen, ylen int) []point {
	points := make([]point, 0, 4)

	if p.c-1 >= 0 {
		points = append(points, point{p.c - 1, p.r})
	}

	if p.c+1 < xlen {
		points = append(points, point{p.c + 1, p.r})
	}

	if p.r-1 >= 0 {
		points = append(points, point{p.c, p.r - 1})
	}

	if p.r+1 < ylen {
		points = append(points, point{p.c, p.r + 1})
	}

	return points
}

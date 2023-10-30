package parkinglot

// every level has several rows with spots
type Row struct {
	Spots []Spot
}

func (r *Row) CanFit(c *Car) bool {
	// check whether the car can fit in the current row
	return len(r.getSpotsFor(c)) == int(c.NumPlacesNeeded)
}

func (r *Row) ParkCar(c *Car) {
	spots := r.getSpotsFor(c)

	for _, i := range spots {
		r.Spots[i].ParkCar(c)
		c.ParkedIn = append(c.ParkedIn, &r.Spots[i])
	}
}

func (r *Row) getSpotsFor(c *Car) []int {
	num_of_places := c.NumPlacesNeeded
	seq := 0
	firstSpot := 0
	for i, spot := range r.Spots {
		if seq == int(num_of_places) {
			break
		}

		if spot.Size >= c.Size {
			seq++
		} else {
			seq = 0
			firstSpot = i + 1
		}
	}

	if seq != int(num_of_places) {
		return []int{}
	}

	spots := make([]int, 0, num_of_places)
	for i := firstSpot; i < firstSpot+int(num_of_places); i++ {
		spots = append(spots, i)
	}

	return spots
}

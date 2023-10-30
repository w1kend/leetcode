package parkinglot

// the hole parking area
type Parking struct {
	Levels          []Level
	PlacesAvailable int
}

// a level of a parking area
type Level struct {
	Rows            []Row
	PlacesAvailable int
}

func (l *Level) CanFit(c *Car) bool {
	for _, row := range l.Rows {
		if row.CanFit(c) {
			return true
		}
	}

	return false
}

func (l *Level) ParkCar(c *Car) bool {
	for _, row := range l.Rows {
		if row.CanFit(c) {
			row.ParkCar(c)
			return true
		}
	}

	return false
}

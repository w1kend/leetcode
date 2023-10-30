package parkinglot

type Spot struct {
	Number      int64
	Size        int8
	IsAvailable bool
	Car         *Car

	Row   *Row
	Level *Level
}

func (s *Spot) Free() {
	s.IsAvailable = true
	s.Car = nil

	s.Level.PlacesAvailable++
}

func (s *Spot) ParkCar(c *Car) {
	// check for availability

	s.IsAvailable = false
	s.Car = c

	s.Level.PlacesAvailable--
}

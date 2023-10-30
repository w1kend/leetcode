package parkinglot

type Car struct {
	Size            int8
	PlateNumber     string
	NumPlacesNeeded int8
	ParkedIn        []*Spot
}

func (c *Car) LeaveParking() {
	for _, spot := range c.ParkedIn {
		spot.Free()
	}
	c.ParkedIn = []*Spot{}
}

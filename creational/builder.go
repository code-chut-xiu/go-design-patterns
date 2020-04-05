package creational

const (
	carWheels, carSeats   int    = 4, 5
	carStructure          string = "Car"
	bikeWheels, bikeSeats int    = 2, 2
	bikeStructure         string = "Bicycle"
	busWheels, busSeats   int    = 10, 30
	busStructure          string = "Bus"
)

//Product
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//Director
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetWheels().SetSeats().SetStructure()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = carWheels
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = carSeats
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = carStructure
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = bikeWheels
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = bikeSeats
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = bikeStructure
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = busWheels
	return b
}

func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = busSeats
	return b
}

func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = busStructure
	return b
}

func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}

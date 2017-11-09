package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return nil
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type ManufactureDirector struct {
	builder BuildProcess
}

func (f *ManufactureDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

func (f *ManufactureDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

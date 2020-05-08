package main

// Builder design pattern tried to
// 1. Abstract complex creations so that object creation is seperated from object user.
// 2. Create an object step by step by filling it's fields and creating embedded objects.
// 3. Reuse the object creation algorithm between many objects.
// Builder design pattern is often described as a relationship between a director and builders.
// Director in charge of construction of objects.
// Builders are the ones that return actual products.

// BuildProcess is a procceding interface defines each steps needed for building vehicle.
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// VehicleProduct is the product that we want to retrieve while using manufacturing.
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// ManufacturingDirector is the one that in charge of accepting builders.
// It has a Construct method that will use builder that is stored in Manufacturing, and will reproduce the required steps.
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

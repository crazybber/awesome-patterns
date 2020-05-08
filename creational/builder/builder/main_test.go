package main

import (
	"testing"
)

func TestBuilderPatter(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()
	if car.Wheels != 4 {
		t.Errorf("number of wheels in car should be 4, it is %d", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	bike := bikeBuilder.GetVehicle()
	if bike.Wheels != 2 {
		t.Errorf("number of wheels on bike should be 2, it is %d", bike.Seats)
	}
}

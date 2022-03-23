package main

import (
	"testing"
)

func TestCarFactory(t *testing.T) {
	carFactory, err := CreateVehicleFactory(CAR)
	if err != nil {
		t.Fatal(err)
	}

	luxuryCar, err := carFactory.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	car, ok := luxuryCar.(Car)
	if !ok {
		t.Fatal("struct assertion failed")
	}

	t.Logf("CarType: Luxury NumWheels=%d, NumSeats=%d, NumDoors=%d",
		luxuryCar.NumWheels(), luxuryCar.NumSeats(), car.NumDoors())
}

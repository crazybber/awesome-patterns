package main

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

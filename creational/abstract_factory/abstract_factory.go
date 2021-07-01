package main

import (
	"errors"
	"fmt"
)

const (
	CAR  = 1
	BIKE = 2
)

func CreateVehicleFactory(v int) (VehicleFactory, error) {
	switch v {
	case CAR:
		return new(CarFactory), nil
	case BIKE:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory of type %d not exist\n", v))
	}
}

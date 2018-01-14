package main

type CruiseMotorbike struct{}

func (s *CruiseMotorbike) NumWheels() int {
	return 2
}

func (s *CruiseMotorbike) NumSeats() int {
	return 2
}

func (s *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}

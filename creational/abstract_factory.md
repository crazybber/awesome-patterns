# Abstract Factory Pattern

Provide an interface for creating families of related or dependent objects without specifying their concrete classes

## Implementation

```go
const (
    CAR  = 1
    LuxuryCarType = 1
    FamilyCarType = 2
)

type Car interface {
    NumDoors() int
}

func CreateVehicleFactory(v int) (VehicleFactory, error) {
    switch v {
    case CAR:
        return new(CarFactory), nil
    default:
        return nil, errors.New(fmt.Sprintf("Factory of type %d not exist\n", v))
    }
}

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
    switch v {
    case LuxuryCarType:
        return new(LuxuryCar), nil
    case FamilyCarType:
        return new(FamilyCar), nil
    default:
        return nil, errors.New(fmt.Sprintf("Vehicle of type %d not exist\n", v))
    }
}

type FamilyCar struct{}

func (*FamilyCar) NumDoors() int {
    return 5
}

func (*FamilyCar) NumWheels() int {
    return 4
}

func (*FamilyCar) NumSeats() int {
    return 5
}

type LuxuryCar struct{}

func (*LuxuryCar) NumDoors() int {
    return 4
}

func (*LuxuryCar) NumWheels() int {
    return 4
}

func (*LuxuryCar) NumSeats() int {
    return 5
}

type Vehicle interface {
    NumWheels() int
    NumSeats() int
}

type VehicleFactory interface {
    NewVehicle(v int) (Vehicle, error)
}
```

## Usage


```go
carFactory, err := CreateVehicleFactory(CAR)
if err != nil {
	panic(err)
}

luxuryCar, err := carFactory.NewVehicle(LuxuryCarType)
if err != nil {
	panic(err)
}

car, ok := luxuryCar.(Car)
if !ok {
	panic(err)
}

fmt.Println(luxuryCar.NumWheels(), luxuryCar.NumSeats(), car.NumDoors())
```

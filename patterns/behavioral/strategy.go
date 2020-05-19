package main

import "fmt"

/*
Code against an interface, not an implementation. Dependency Injection! :)
*/

func main() {
	var driver Driver

	driver = NewDriver(PersonalCar{})
	driver.vehicle.DriveToTheAirport()

	driver = NewDriver(ShuttleBus{})
	driver.vehicle.DriveToTheAirport()
}

type Vehicle interface {
	DriveToTheAirport()
}

type Driver struct {
	vehicle Vehicle
}

func NewDriver(vehicle Vehicle) Driver {
	return Driver{
		vehicle: vehicle,
	}
}

type PersonalCar struct {
}

func (p PersonalCar) DriveToTheAirport() {
	fmt.Println("Driving in my own car")
}

type ShuttleBus struct {
}

func (s ShuttleBus) DriveToTheAirport() {
	fmt.Println("Taking the shuttle")
}

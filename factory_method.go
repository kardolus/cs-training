package main

import (
	"fmt"
)

/*
	A factory method gives you an implmentation of an interface. Usually the input is a
	constant. Sometimes constructed objects are reused instead of re-instantiated.
*/

const (
	SportsCar = 1 << iota
	FamilyCar
	Truck
)

func main() {
	var car Car

	car, _ = FactoryMethod(SportsCar)
	car.Drive()

	car, _ = FactoryMethod(FamilyCar)
	car.Drive()

	car, _ = FactoryMethod(Truck)
	car.Drive()
}

type Car interface {
	Drive()
}

func FactoryMethod(carType int) (Car, error) {
	switch carType {
	case SportsCar:
		return Tesla{}, nil
	case FamilyCar:
		return Honda{}, nil
	case Truck:
		return Uhaul{}, nil
	}

	return nil, fmt.Errorf("car not found!")
}

type Tesla struct {
}

func (t Tesla) Drive() {
	fmt.Println("Driving a hella fast Tesla")
}

type Uhaul struct {
}

func (u Uhaul) Drive() {
	fmt.Println("Moving stuff in my Uhaul")
}

type Honda struct {
}

func (h Honda) Drive() {
	fmt.Println("Taking the family on a drive")
}

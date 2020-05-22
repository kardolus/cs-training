package main

import "fmt"

/*
	Convert the interface of a class into another interface clients expect. Adapter lets
	classes work together that couldn't otherwise because of incompatible interfaces.
*/

// Just some pseudo code to get the idea

type EuroPlug interface {
	PlugIt()
}

type AmericanPlug interface {
	PlugIt()
}

type EuroShaver struct {
}

func (e EuroShaver) PlugIt() {
	fmt.Println("Plugging in a European Shaver")
}

type USShaver struct {
}

func (u USShaver) PlugIt() {
	fmt.Println("Plugging in a US Shaver")
}

func Adapt(europlug EuroPlug) AmericanPlug {
	var result AmericanPlug

	// set fields on the American Plug based on the Euro Plug

	return result
}

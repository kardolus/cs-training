package main

import (
  "fmt"
)

/*
  Think about a happy meal :). It consist of a main item, a side item, a drink, 
  and a toy. The main item can be a chicken burger or a veggie burger or whatever.
  The drink can be a coke or a sprite etc. But the process of constructing a happy
  meal always consists of the same steps.
*/

func main() {  
  fmt.Println("Make me a happy meal!")
  
  var assembly HappyMeal
  assembly = assembly.Side("Fries").Drink("Coke").(HappyMeal)
  
  fmt.Printf("Assembly: %+v\n", assembly)
  
  var choiceOne = assembly.Main("Veggie").Toy("Motorcycle").Build()
  var choiceTwo = assembly.Main("BatBurder").Toy("Helicopter").Build()
  
  fmt.Printf("%+v\n", choiceOne)
  fmt.Printf("%+v\n", choiceTwo)
 
  choiceOne.Eat()
  choiceTwo.Eat()
}

type Dinner interface {
  Eat() 
}

type Builder interface {
  Build() Dinner
  Main(choice string) Builder
  Side(choice string) Builder
  Drink(choice string) Builder
  Toy(choice string) Builder
}

type SomeMeal struct {
  YourMeal string
}

func (s SomeMeal) Eat() {
  fmt.Println("Cook me some food")
}

type HappyMeal struct {
  YourMain string
  YourSide string
  YourDrink string
  YourToy string
}

func (h HappyMeal) Main(choice string) Builder {
  fmt.Println("Your pick:", choice)
  
  h.YourMain = choice
  
  return h
}

func (h HappyMeal) Side(choice string) Builder {
  fmt.Println("Your pick:", choice)
  h.YourSide = choice
  return h
}

func (h HappyMeal) Drink(choice string) Builder {
  fmt.Println("Your pick:", choice)
  h.YourDrink = choice
  return h
}

func (h HappyMeal) Toy(choice string) Builder {
  fmt.Println("Your pick:", choice)
  h.YourToy = choice
  return h
}

func (h HappyMeal) Build() Dinner {
  var meal = fmt.Sprintf("%s %s %s %s", h.YourMain, h.YourSide, h.YourDrink, h.YourToy)
  return SomeMeal{YourMeal: meal}
}



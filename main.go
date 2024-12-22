// ВАРИАНТ кода без DI
// При такой реализации сложно давить новый тип печки (тип ovenType) в существующую структуру BakeryNoDI
package main

import (
	"fmt"
)

//region Необходимые интерфейсы и структуры для реализации DI

//region Интерфейсы и методы в них

type Oven interface {
	Heat() string
}

type Ingredient interface {
	Mix() string
}

//endregion

//region Структуры и реализации методов интерфейсов

type GasOven struct{}

func (o GasOven) Heat() string {
	return "Heating with gas oven!"
}

type ElectricOven struct{}

func (o ElectricOven) Heat() string {
	return "Heating with electric oven!"
}

type Sugar struct{}
type Flour struct{}
type Butter struct{}

func (i Sugar) Mix() string {
	return "Mixing sugar"
}

func (i Flour) Mix() string {
	return "Mixing with flour"
}
func (i Butter) Mix() string {
	return "Adding butter"
}

//endregion

//endregion

type Bakery struct {
	oven        Oven
	ingredients []Ingredient
}

func (b *Bakery) Bake() {
	fmt.Println(b.oven.Heat())
	for _, ingredient := range b.ingredients {
		fmt.Println(ingredient.Mix())
	}

	fmt.Println("Baking an awesome pie!")
}

func main() {
	gasOven := GasOven{}
	electricOven := ElectricOven{}
	sugar := Sugar{}
	butter := Butter{}
	flour := Flour{}

	bakery := Bakery{oven: gasOven, ingredients: []Ingredient{flour, sugar}}
	bakery.Bake()

	bakery = Bakery{oven: electricOven, ingredients: []Ingredient{sugar, butter}}
	bakery.Bake()

}

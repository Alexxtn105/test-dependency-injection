// ВАРИАНТ кода без DI

package main

import "fmt"

type Bakery struct {
	ovenType    string
	ingredients []string
}

func (b *Bakery) Bake() {
	switch b.ovenType {
	case "gas oven":
		fmt.Println("Heating with gas oven!")

	case "electric oven":
		fmt.Println("Heating with gas oven!")
	}
	for _, ingredient := range b.ingredients {
		switch ingredient {
		case "flour":
			fmt.Println("Mixing with flour")
		case "sugar":
			fmt.Println("Mixing sugar")
		case "butter":
			fmt.Println("Adding butter")

		}
	}

	fmt.Println("Baking an awesome pie!")
}

func main() {
	bakery := Bakery{ovenType: "gas oven", ingredients: []string{"flour", "sugar"}}
	bakery.Bake()

	bakery = Bakery{ovenType: "electric oven", ingredients: []string{"sugar", "butter"}}
	bakery.Bake()

}

// ВАРИАНТ кода без DI
// При такой реализации сложно давить новый тип печки (тип ovenType) в существующую структуру BakeryNoDI
package main

import "fmt"

type BakeryNoDI struct {
	ovenType    string
	ingredients []string
}

func (b *BakeryNoDI) Bake() {
	switch b.ovenType {
	case "gas oven":
		fmt.Println("Heating with gas oven!")

	case "electric oven":
		fmt.Println("Heating with electric oven!")
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
	bakery := BakeryNoDI{ovenType: "gas oven", ingredients: []string{"flour", "sugar"}}
	bakery.Bake()

	bakery = BakeryNoDI{ovenType: "electric oven", ingredients: []string{"sugar", "butter"}}
	bakery.Bake()

}

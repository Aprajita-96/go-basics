package main

import (
	"fmt"

	utils "example.com/utils"
)

func main() {

	products := []*utils.Product{
		&Product{id: 5, name: "Pen", cost: 10, units: 100},
		&Product{id: 7, name: "Pencil", cost: 5, units: 50},
		&Product{id: 2, name: "Marker", cost: 50, units: 30},
		&Product{id: 6, name: "Scribble Pad", cost: 30, units: 40},
		&Product{id: 9, name: "Mouse", cost: 100, units: 10},
	}
	fmt.Println("Products")
	print(products)
	fmt.Println()

	fmt.Println("Under stocked products [units < 40]")
	underStockedProductsPredicate := func(p *utils.Product) bool {
		return p.units < 40
	}
	underStockedProducts := utils.FilterProducts(products, underStockedProductsPredicate)
	print(underStockedProducts)
	fmt.Println()

	fmt.Println("Well stocked products [!understocked]")
	/*
		wellStockedProductsPredicate := func(p *Product) bool{
			return !underStockedProductsPredicate(p);
		}
	*/
	wellStockedProductsPredicate := utils.Negate(underStockedProductsPredicate)
	wellStockedProducts := utils.FilterProducts(products, wellStockedProductsPredicate)
	print(wellStockedProducts)
	fmt.Println()

	fmt.Println("Costly products [cost >= 50]")
	costlyProductPredicate := func(p *utils.Product) bool {
		return p.cost >= 50
	}
	costlyProducts := utils.FilterProducts(products, costlyProductPredicate)
	print(costlyProducts)

	fmt.Println()
	fmt.Println("Cheaper Products [!costly]")
	/*
		cheaperProductPredicate := func(p *Product) bool{
			return !costlyProductPredicate(p)
		};
	*/
	cheaperProductPredicate := utils.Negate(costlyProductPredicate)
	cheaperProducts := utils.FilterProducts(products, cheaperProductPredicate)
	print(cheaperProducts)

	cheaperAndWellStockedPredicate = utils.And(cheaperProductPredicate, wellStockedProductsPredicate)

}

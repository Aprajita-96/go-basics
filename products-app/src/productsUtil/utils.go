package utils

import "fmt"

type Product struct {
	id    int
	name  string
	cost  float64
	units int
}

type Predicate func(p *Product) bool

func FilterProducts(products []*Product, predicate Predicate) []*Product {
	var result = []*Product{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func filterCostlyProducts(products []*Product) []*Product {
	//product.cost >= 50
	var result = []*Product{}
	for _, product := range products {
		if product.cost >= 50 {
			result = append(result, product)
		}
	}
	return result
}

func filterUnderStockedProducts(products []*Product) []*Product {
	//product.units < 40
	var result = []*Product{}
	for _, product := range products {
		if product.units < 40 {
			result = append(result, product)
		}
	}
	return result
}

func Print(products []*Product) {
	ForEach(products, func(product *Product) {
		fmt.Println(product.id, product.name, product.cost, product.units)
	})
}

func ForEach(products []*Product, action func(p *Product)) {
	for _, product := range products {
		action(product)
	}
	return
}

func Negate(predicate Predicate) Predicate {
	return func(p *Product) bool {
		return !predicate(p)
	}
}

func Or(predicate1 Predicate, predicate2 Predicate) Predicate {
	return func(p *Product) bool {
		return predicate1(p) || predicate2(p)
	}
}

func And(predicate1 Predicate, predicate2 Predicate) Predicate {
	return func(p *Product) bool {
		return predicate1(p) && predicate2(p)
	}
}

package main

import (
	"fmt"
	"math"
)

var resultrectangle float64
var resultcircle float64

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x, y, r float64
}
type Triangle struct {
	x, y, z float64
}

type Shape interface {
	area() float64
}

func (tri *Triangle) area() float64 {
	return float64(30)
}

func (rect *Rectangle) area() float64 {
	x := rect.x2 - rect.x1
	y := rect.y2 - rect.y1
	return math.Sqrt(x*x + y*y)
}

func (circle *Circle) area() float64 {
	return math.Pi * circle.r * circle.r
}
func SumOfAreas(rectangle []*Rectangle, circle []*Circle) (sum float64) {
	var result float64
	for _, rec := range rectangle {
		result += rec.area()
	}
	for _, cir := range circle {
		result += cir.area()
	}
	fmt.Println(result)
	return result
}

func main() {
	rectangle := Rectangle{x1: 10, x2: 20, y1: 10, y2: 20}
	rectangle1 := Rectangle{x1: 30, x2: 20, y1: 10, y2: 30}
	rectangle2 := Rectangle{x1: 40, x2: 20, y1: 40, y2: 30}
	rectangle3 := Rectangle{x1: 20, x2: 20, y1: 40, y2: 30}
	var collectionOfrectangles = []*Rectangle{&rectangle, &rectangle1, &rectangle2, &rectangle3}
	circle := &Circle{x: 10, y: 20, r: 10}
	circle1 := &Circle{x: 10, y: 20, r: 20}
	circle2 := &Circle{x: 10, y: 20, r: 30}
	var collectionOfcircles = []*Circle{circle, circle1, circle2}
	SumOfAreas(collectionOfrectangles, collectionOfcircles)

	//open close principle , extend the above code without modifying it.
	//ex:sumOf areas Function
	//use case :Add another shape: Triangle
	var collectionOfTri = []*Triangle{
		&Triangle{x: 10, y: 20, z: 30},
		&Triangle{x: 10, y: 20, z: 30},
	}

	shapes := []Shape{
		&Rectangle{x1: 10, x2: 20, y1: 20, y2: 40},
		&Rectangle{x1: 10, x2: 20, y1: 20, y2: 40},
		&Circle{x: 10, y: 20, r: 10},
		&Circle{x: 10, y: 20, r: 20},
		&Circle{x: 10, y: 20, r: 30},
		&Triangle{x: 10, y: 20, z: 30},
		&Triangle{x: 10, y: 20, z: 30},
		&Triangle{x: 10, y: 20, z: 30},
	}
	fmt.Println(sumOfAreasOpenClosePrinciple(shapes))

}
func sumOfAreasOpenClosePrinciple(shapes []Shape) float64 {
	sumOfArea := float64(0)
	for _, shape := range shapes {
		sumOfArea += shape.area()
	}
	return sumOfArea
}

/*
to do:
	create a collection of rectangles
	create a collection of circles
	create a function 'sumOfAreas()' which will calculate the sum of areas of rectangles and circles
*/

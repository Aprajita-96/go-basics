package main

import "fmt"

type rect struct {
	width, height float64
}

func (r *rect) areas() float64 {
	return r.width * r.height
}
func (r rect) perimm() float64 {
	return 2*r.width + 2*r.height
}

func methodmain() {
	//Go automatically handles conversion between values and
	//pointers for method calls.
	//You may want to use a pointer receiver type to avoid copying
	//on method calls or to allow the method to mutate the receiving
	//struct.
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.areas())
	fmt.Println("perim:", r.perimm())

	rp := &r
	fmt.Println("area: ", rp.areas())
	fmt.Println("perim:", rp.perimm())
}

package main

import "fmt"

type Employee struct {
	age    int
	name   string
	salary float32
}

//make is used for slices arrays and map
//new is for struct , returns a pointer
//update as a func
func update(e *Employee) {
	e.age = 89
	fmt.Println(e)
}

//make a method for struct : this is update as a method
func (e *Employee) update() {
	e.age = 89
	fmt.Println(e)
}

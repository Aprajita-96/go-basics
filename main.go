package main

import (
	"fmt"
)

func init() {
	fmt.Print("init triggered")
}

func mainfun() {
	fmt.Println("hello world")
	//thid is a variable
	var a = "intial variable"
	fmt.Print(a)

	//short hand way, also a constant and cannot be changed anytime in program
	apple := "apple"
	fmt.Println(apple)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	//will print 0 because int is intialised with 0
	fmt.Println(e)

}

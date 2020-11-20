package main

import "fmt"

//when u declare variables outside the function , declare them with var
//we define them outside the function when we want our own type
var a int = 10

func main() {
	//way1
	var x int
	x = 10
	fmt.Println(x)

	//way2
	var y int = 10
	fmt.Println(y)

	//way3
	z := 10
	fmt.Println(z)

	//way 4
	var l, m = 10, 20
	fmt.Println(l + m)

	//way 5
	var (
		del int = 10
		dey     = 20
	)
	fmt.Println(del + dey)

	const (
		constantvariable = 10
	)

}

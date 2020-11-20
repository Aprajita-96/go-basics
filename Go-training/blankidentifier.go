package main

import (
	"fmt"
)

func main() {

	fmt.Println(add(3, 4))
	fmt.Println(divide(100, 100))
	var result, remainder = divide(100, 100)
	fmt.Printf("%v, %v\n", result, remainder)
	//if u dont want to print the remainder
	var res, _ = divide(100, 100)
	fmt.Printf("%v\n", res)
}
func add(a, b int) (result int) {
	return a + b
}

func divide(x, y int) (int, int) {
	result := x / y
	remainder := x % y
	return result, remainder

}

package main

import "fmt"

//anonymous functions are useful wjhen we want to define a function inline without having to name it
func intsqence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closuremain() {

	nextInt := intsqence()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}

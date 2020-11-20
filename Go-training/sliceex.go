package main

import (
	"fmt"
)

func main() {
	//auto initialised with the number of numbers u give no size, []
	nos := []int{3, 1, 2, 4, 5, 7} // using the slice literal
	// nos:= make ([] int, 6)
	// nos[1]=1
	fmt.Println(nos)
	fmt.Printf("%v\n", nos)

	newNos := nos[1:4]
	fmt.Printf("%v\n", newNos)

	oneNo := nos[1:2] //1,(2-1)
	fmt.Printf("%v\n", oneNo)

	//making slice dynamic?
	nos_dynamic := []int{}
	nos_dynamic = append(nos, 3, 1, 2, 3, 7)
	nos_dynamic = append(nos, 1)
	fmt.Println(nos_dynamic)
	//append two slices
	nos1 := []int{1, 2, 3}
	nos2 := []int{4, 5, 6}
	//append one slice into another
	newNum := append(nos1, nos2...)
	fmt.Println(newNum)

	// append(nos1,nos2) will throw error :./sliceex.go:30:18: cannot use nos2 (type []int) as type int in append
	// which means that nos1 is a slice of type integers so it cannot append a slice inside it, it would work if we say
	// nos1:= [][]int{} which means it is a slice of type slice

	for index, value := range nos {
		fmt.Println(value)
		if index%4 == 0 {
			break
		}
	}

	//zero value of a slice is nil , if it is not initialised
	var empty []int
	empty = append(empty, 10, 20)
	fmt.Println(empty)
	// dummy := []int{} //initialised
	if empty == nil {
		fmt.Println("dummy is empty")
	}
}

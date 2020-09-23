package main

import "fmt"

func zerovalue(a int) {
	a = 0
}
func pointervalue(ptr *int) {
	*ptr = 0
}

func pointemain() {

	i := 1
	fmt.Println("initial:", i)

	zerovalue(i)
	fmt.Println("not changes value", i)

	pointervalue(&i)
	fmt.Println("changed value", i)

}

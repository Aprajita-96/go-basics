package main

import "fmt"

func switchmain() {

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Im a boolean")
		case int:
			fmt.Println("im an integer")
		default:
			fmt.Printf("Dont know the type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Aparajita")

}

package main

import "fmt"

type Person struct {
	id   int
	name string
}

// type Employe struct {
// 	id     int
// 	name   string
// 	salary float64
// }
//composition
type Employe struct {
	Person
	salary float64
}

func main() {

	emp := Employe{Person: Person{id: 100, name: "Apps"}, salary: 1000}
	empp := Employe{}
	empp.id = 1000
	empp.name = "Appu"
	empp.salary = 900

	fmt.Println(emp)
	fmt.Println(empp)
}

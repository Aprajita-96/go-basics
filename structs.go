package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 52
	return &p
}

func structsmain() {
	fmt.Println(person{"Bob", 2})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Appu"))

	s := person{name: "Annu", age: 90}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
}

package main

import (
	"fmt"
)

func main() {

	// emp:= Employee {age:20,name:"Aprajita",Salary:8000}
	emp := new(Employee)
	emp.name = "Apps"
	emp.salary = 0
	fmt.Println(emp)
	fmt.Println(emp.age)
	emp.update()
	update(emp)
}

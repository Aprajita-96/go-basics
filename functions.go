package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusthree(a, b, c int) int {
	return a + b + c
}

//a function can also return multiple values like
func values() (int, int) {
	return 3, 7
}

//a function can take in multiple values
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func functiomain() {

	res := plus(1, 2)
	fmt.Println("1+2: ", res)

	result := plusthree(1, 2, 3)
	fmt.Println("1+2+3:", result)

	a, b := values()
	fmt.Println(a, b)

	//u can pass any number of integers to sum
	sum(1, 2)
	sum(1)
	sum(1, 2, 3, 4)

	arrayofnumbers := []int{1, 2, 3, 4, 5, 6, 7}
	sum(arrayofnumbers...)

}

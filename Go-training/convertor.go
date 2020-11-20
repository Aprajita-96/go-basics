package main

import (
	"fmt"
)

func main() {
	var (
		output        float64
		input         float64
		optionChoosed int
	)
	fmt.Println("Please choose the number:")
	fmt.Println("1.Fahrenheit to celsius")
	fmt.Println("2.Feet to metres")
	fmt.Scanf("%d", &optionChoosed)
	fmt.Println("enter the value:")
	fmt.Scanf("%f", &input)
	if optionChoosed == 1 {
		output = (input - 32) * 5 / 9
		fmt.Printf("Temperature in celcius is: %f ", output)
	} else {
		output = (input * (0.348))
		fmt.Print("Value in metres is: %f", output)
	}
}

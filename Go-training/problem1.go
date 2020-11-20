/*
organize each name within a slice based on its length and
print the names of the given length (taken as input from the user)
*/
package main

import (
	"fmt"
)

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}
var namesearch = map[int][]string{
	1: []string{},
	2: []string{},
	3: []string{"Mia", "Ava"},
	4: []string{"Evan", "Neil", "Adam", "Matt", "Emma"},
	5: []string{"Martin", "Emily", "Chloe"},
	6: []string{"Olivia", "Sophia", "Alexis"},
	7: []string{"Katrina", "Madison", "Abigail", "Addison", "Natalie"},
	8: []string{"Isabella", "Samantha"},
	9: []string{"Elizabeth"},
}

func main() {
	// insert your code here
	namesMap := map[int][]string{}
	//dynamic map creation
	for _, value := range names {
		namesMap[len(value)] = append(namesMap[len(value)], value)
	}
	fmt.Println("Enter the length of alphabet:")
	for {
		output := givenames()
		if output == 0 {
			break
		}
	}
}
func givenames() (out int) {
	var input int
	fmt.Scanln(&input)
	if input == 0 {
		fmt.Println("invalid input")
		return 0
	}
	switch input {
	case 1:
		fmt.Println(namesearch[1])
	case 2:
		fmt.Println(namesearch[2])
	case 3:
		fmt.Println(namesearch[3])
	case 4:
		fmt.Println(namesearch[4])
	case 5:
		fmt.Println(namesearch[5])
	case 6:
		fmt.Println(namesearch[6])
	case 7:
		fmt.Println(namesearch[7])
	case 8:
		fmt.Println(namesearch[8])
	case 9:
		fmt.Println(namesearch[9])
	default:
		fmt.Println("Please enter valid number")
	}
	return 1
}

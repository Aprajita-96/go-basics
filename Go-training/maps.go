package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 0
	m["k2"] = 13

	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("v1", v1)

	fmt.Println("len", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//the below output is printed as zero which makes it unclear that whether the k2 is present or the value of k2 is 0
	fmt.Println("map:", m["k2"])
	//hence we use _, to check whether the key is present or not, below output is false bevcause key is not present
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("intialised map:", n)

	states := map[string]string{
		"KA": "Bengaluru",
		"TN": "Chennai",
		"TS": "Hyderabad",
	}

	fmt.Println(states)

	states["OD"] = "Bhuvaneshwar"

	delete(states, "TN")

	//check if key exists
	capital, ok := states["DL"]
	fmt.Println(capital, ok)

	if _, exists := states["OD"]; exists {
		fmt.Println("OD exists")
	} else {
		fmt.Println("Doesnt exists")
	}

}

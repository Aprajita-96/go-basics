package main

import "fmt"

//range iterates over elements in a variety of ds.
func rangemain() {
	//example of arrays iteration with and without index
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index of 3:", i)
		}
	}

	//exmaple of maps
	mapsexample := map[string]string{"a": "apple", "b": "bat", "c": "cat"}

	for key, value := range mapsexample {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range mapsexample {
		fmt.Println("keys:", key)
	}

}

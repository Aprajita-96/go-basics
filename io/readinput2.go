// read input from the console:
package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var inputt string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin) // reader for input
	fmt.Println("Please enter some input: ")
	inputt, err = inputReader.ReadString('\n')

	if err == nil {
		fmt.Printf("The input was: %s\n", inputt)
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// var inputFile *os.File
	// var inputError, readerError os.Error
	// var inputReader *bufio.Reader
	// var inputString string
	var (
		avgsalary float64
		format    = "%d ,%s , %f"
		id        int
		name      string
		salary    float64
		count     int
		total     float64
	)

	inputFile, inputError := os.Open("emp.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break // error or EOF
		}
		fmt.Printf("The input was: %s", inputString)
		fmt.Sscanf(inputString, format, &id, &name, &salary)
		fmt.Println(id, name, salary)
		count++
		total += salary
	}
	avgsalary += float64(total) / float64(count)
	fmt.Println(avgsalary)
}

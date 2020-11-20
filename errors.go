package main

import (
	"errors"
	"fmt"
)

//1. errors are by convention the last return value
//2. errors.new constructs a new error value with the given message
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant work with 42")
	}

	// a nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

// example 2 . says and depicts the custom errors(like the custom error class which is made in java)

// so in place of a class , we make a struct
type argError struct {
	arg  int
	prob string
}

//receives a
func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func errorsmain() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

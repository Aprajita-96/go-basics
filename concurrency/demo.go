package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}
var data = "Hello"

// func say(str string, x chan int, y chan int) {
// 	for i := 0; ; i++ {
// 		<-x
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println(str)
// 		y <- i
// 	}
// }

func main() {
	// c1 := make(chan int)
	// c2 := make(chan int)

	// go say("Hello", c1, c2)
	// go say("World", c2, c1)

	// c1 <- -1
	// var input string
	// fmt.Scanln(&input)

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		// for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(data)
		data = "World"

		// }
	}()
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		// for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(data)
		data = "Hello"
		// }
	}()

	var input string
	fmt.Scanln(&input)
}

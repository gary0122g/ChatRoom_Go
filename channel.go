package main

import (
	"fmt"
	"time"
)

func SendValues(c chan string) {
	fmt.Println("Executing Goroutine")
	time.Sleep(1 * time.Second)
	c <- "8"
	fmt.Println("Finished")
}

// func main() {
// 	values := make(chan string, 2)
// 	defer close(values)

// 	go SendValues(values)
// 	go SendValues(values)
// 	value := <-values
// 	fmt.Println(value)
// }

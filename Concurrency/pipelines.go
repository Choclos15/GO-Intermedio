package main

import "fmt"

func Generator(c chan<- int) {
	for value := 1; value <= 10; value++ {
		c <- value
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func Print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

// func main() {
// 	generator := make(chan int)
// 	doubles := make(chan int)

// 	go Generator(generator)
// 	go Double(generator, doubles)
// 	Print(doubles)
// }

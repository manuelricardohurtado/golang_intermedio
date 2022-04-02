package main

import "fmt"

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c) //Cerrar canal C
}

func Doble(in <-chan int, out chan<- int) {
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

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Doble(generator, doubles)
	Print(doubles)
}

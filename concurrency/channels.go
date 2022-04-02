package main

import "fmt"

func main() {
	c := make(chan int, 3) // en este caso 3 es el maximo buffer del canal

	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)

}

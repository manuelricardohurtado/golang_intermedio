package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int { //Función Anonima
		return x * 2
	}()
	fmt.Println(y)

	c := make(chan int)
	go func() { //Función Aninima
		fmt.Println("Starting Function")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}

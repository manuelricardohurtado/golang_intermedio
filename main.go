package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//Delcaración de Variables
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)
	//Manejo de Errores
	myValue, err := strconv.ParseInt("21", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	//Mapas (LLave - Valor)
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])

	//Slice
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	//Agregar elementos al final del Slice
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	//llamado de GoRutine
	c := make(chan int)
	go doSometing(c)
	<-c

	//Puntero
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

//Funcion con GoRutine
func doSometing(c chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("DONE¡¡¡")
	c <- 1
}

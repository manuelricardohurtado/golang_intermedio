package main

import "fmt"

func sum(values ...int) int { // Función que recibe  varios valores como Slice
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) { //Función que recibe varios nombres como Slice y los imprime
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quad int) { //Funciónes Variádicas retorna todas la variables de salida con un solo Return
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	printNames("Manuel", "Ricardo", "Catherine", "Isabel")
	fmt.Println(getValues(3))
}

package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

//Funcion constructora para la Forma 4
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//Forma Numero 1:
	e := Employee{}
	fmt.Printf("%v\n", e)
	//Forma Numero 2:
	e2 := Employee{
		id:       2,
		name:     "Name 2",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	//Forma Numero 3:
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 3
	e3.name = "Name 3"
	e3.vacation = true
	fmt.Printf("%v\n", *e3)

	//Forma Numero 4 (Forma Sugerida):
	e4 := NewEmployee(4, "Name 4", true)
	fmt.Println(*e4)

	//Punteros
	a := 10
	b := &a
	c := *b
	fmt.Println(a, b, c)
}

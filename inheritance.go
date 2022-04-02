package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessge() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessge() string {
	return "Temporary Time Employee"
}

type PrintInfo interface {
	getMessge() string
}

func getMessge(p PrintInfo) {
	fmt.Println(p.getMessge())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Ivanchis"
	ftEmployee.age = 33
	ftEmployee.id = 555
	fmt.Println(ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessge(tEmployee)
	getMessge(ftEmployee)
}

package main

import "fmt"

type contactInfo struct {
	email string
	zipcode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// First way of defining structs
	//alex := person{"Alex", "Lanier"}
	// Second way
	// alex := person{firstName: "Alex", lastName: "Lanier"}
	// Third way
	// var alex person
	// fmt.Printf("%+v", alex)
	// alex.firstName = "Alex"
	// alex.lastName = "Lanier"
	// fmt.Printf("%+v", alex)
	alex := person{
		firstName: "Alex",
		lastName: "Lanier",
		contact: contactInfo{
			email: "what@what.com",
			zipcode: 87541,
		},
	}
	fmt.Printf("%+v", alex)
	fmt.Println("Second Print")
	alex.print()
	alex.updateName("John");
	alex.print()
	fmt.Println("\n Print using pointer")
	greatProgram()
}

// adding receiver function

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(name string) {
	p.firstName = name;
}

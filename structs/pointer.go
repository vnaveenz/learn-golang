package main

import "fmt"

type greatContactInfo struct {
	email string
	zipcode int
}

type greatPerson struct {
	firstName string
	lastName string
	contact greatContactInfo
}

func (gpp *greatPerson) greatPersonUpdate(newName string) {
	(*gpp).firstName = newName
}

func (p greatPerson) printGreatPerson() {
	fmt.Printf("%+v", p)
}

func greatProgram() {
	nav := greatPerson {
		firstName: "Nav",
		lastName: "ven",
		contact: greatContactInfo{
			email: "what@what.com",
			zipcode: 87541,
		},
	}
	// OPTION 1
	navPtr := &nav
	navPtr.greatPersonUpdate("Naveen")
	// OPTION 2
	nav.printGreatPerson()
	nav.greatPersonUpdate("Naveeeeeen")
	nav.printGreatPerson()
}

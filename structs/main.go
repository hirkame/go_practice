package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	hiroki := person{
		firstName: "Hiroki",
		lastName:  "Kameyama",
		contact: contact{
			email:   "kame@gm.com",
			zipCode: 111111,
		},
	}

	hiroki.updateName("Hiro")
	hiroki.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

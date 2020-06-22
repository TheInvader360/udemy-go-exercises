package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	j := person{
		firstName: "John",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "john@smith.com",
			zipCode: 12345,
		},
	}
	j.updateName("Jon")
	j.print()
}

func (personPtr *person) updateName(newFirstName string) {
	(*personPtr).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

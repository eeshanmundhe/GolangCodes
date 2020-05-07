package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   []string
}

func (p *Person) getFullName() string {

	//fmt.Println(p.firstName + p.lastName)
	return p.firstName + p.lastName
}

func main() {
	person := &Person{
		firstName: "Eeshan",
		lastName:  "Mundhe",
		age:       21,
		address:   []string{"Mumbai", "Pune", "Thane"},
	}
	//getFullName(&person)
	//variable: = &person
	fmt.Println(person.getFullName())
}

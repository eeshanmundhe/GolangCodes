package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func modifyPointer(p *Person) {
	p.firstName = "Charlie"
	p.lastName = "Harper"
}

func modifyValue(sub Person) {
	sub.firstName = "Eeshan"
	sub.lastName = "Mundhe"
	fmt.Println(sub)
}

func main() {
	person := Person{
		firstName: "Alan",
		lastName:  "Harper",
	}
	fmt.Println(person) //original value printed
	modifyValue(person)
	fmt.Println(person) //again original as variable sub is modified, not person
	modifyPointer(&person)
	fmt.Println(person) //new values printed
}

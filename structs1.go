package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int8
	address   []string
}

func struct_creator(first string, last string, a int8, add []string) person {
	var person3 person
	person3.firstName = first
	person3.lastName = last
	person3.age = a
	person3.address = add
	return person3
}

func main() {
	var details person
	details.firstName = "Eeshan"
	details.lastName = "Mundhe"
	details.age = 21
	details.address = []string{"Mumbai", "Pune"}

	person1 := person{
		firstName: "Ash",
		lastName:  "Ketchum",
		age:       10,
		address:   []string{"Pallet Town", "Kanto"},
	}

	var person2 person
	person2.firstName = "Roman"
	person2.lastName = "Reigns"
	person2.age = 35
	person2.address = []string{"Pensacola", "Florida"}

	fmt.Println(details)
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(struct_creator("ABCD", "WXYZ", 26, []string{"Alphabets", "English"}))
}

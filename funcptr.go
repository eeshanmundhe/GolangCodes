package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   []string
}

func printInfo(p *Person) {
	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
	fmt.Println(p.age)

	for i := 0; i < len(p.address); i++ {
		fmt.Println(p.address[i])
	}
}

func main() {
	person := Person{
		firstName: "Eeshan",
		lastName:  "Mundhe",
		age:       21,
		address:   []string{"Mumbai", "Pune", "Thane"},
	}
	printInfo(&person)
}

package main

import "fmt"

func main() {
	var pet struct {
		name string
		age  int
	}
	pet.name = "Tommy"
	pet.age = 3
	fmt.Println("Name:", pet.name)
	fmt.Println("Age:", pet.age)
}

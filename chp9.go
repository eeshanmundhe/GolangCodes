package main

import "fmt"

type myType string

func (m myType) sayHi() {
	fmt.Println("Hi from", m)
}

func main() {
	value := myType("myTye Value")
	value.sayHi()
	anotherValue := myType("another mytype value")
	anotherValue.sayHi()
}

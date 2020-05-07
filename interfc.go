package main

import "fmt"

type newInterface interface {
	noParmameter()
	oneParmameter(float64)
	methodReturn() string
}
type myType int

func main() {
	var value newInterface
	value = myType(5)
	value.noParmameter()
	value.oneParmameter(127.3)
	fmt.Println(value.methodReturn())
}
func (m myType) noParmameter() {
	fmt.Println("This method has no parameters!")
}
func (m myType) oneParmameter(float64) {
	fmt.Println("This method has one parameter!")
}
func (m myType) methodReturn() string {
	return "This method returned a string"
}

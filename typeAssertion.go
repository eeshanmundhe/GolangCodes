package main

import "fmt"

func main() {
	var i interface{} = 10
	storeString, ok := i.(int)
	fmt.Println(storeString, ok)
	//var num float32 = 17
	checkType(true)
}

func checkType(anyType interface{}) {
	switch anyType.(type) {
	case string:
		fmt.Println("It's a string")

	case float64:
		fmt.Println("It's a float64")

	case float32:
		fmt.Println("It's a float32")

	case int:
		fmt.Println("It's an int")

	default:
		fmt.Println("Other case")
	}
}

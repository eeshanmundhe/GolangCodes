package main

import "fmt"

func main() {
	var a = 25.0
	var b = 10.4
	fmt.Println("Numbers are ", a, "and", b)
	fmt.Println("Sum is ", add(a, b))
	fmt.Println("Difference is ", diff(a, b))
	fmt.Println("Product is ", prod(a, b))
	fmt.Println("Division is ", divide(a, b))
}

func add(x, y float64) float64 {
	return x + y
}

func diff(x, y float64) float64 {
	return x - y
}

func prod(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}

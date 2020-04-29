package main

import "fmt"

func main() {
	result := mathOperations(add)
	fmt.Println(result)
	fmt.Println(mathOperations(sub))
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mathOperations(funcName func(num1, num2 int) int) int {
	result := funcName(50, 100)
	return result
}

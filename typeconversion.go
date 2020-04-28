package main

import "fmt"

func main() {
	var a = 25
	var b = 10.4
	ans := float64(a) * b
	fmt.Println("Int a = 25 multiplied by Float b = 10.4 is", ans)
	var c = 25.0
	var d = 10
	answer := int(c) * d
	fmt.Println("Float c = 25.0 multiplied by int d = 10 is", answer)
}

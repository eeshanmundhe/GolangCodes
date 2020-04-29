package main

import "fmt"

func main() {
	fmt.Println("Printing values greater than 50")
	arr := [10]int{50, 75, 90, 100, 25, 126, 59, 0, 89, 102}
	for i, val := range arr {
		if val > 50 {
			fmt.Println(val, "at position", i, "is grater than 50")
		}
	}
}

package main

import "fmt"

func main() {
	arr := [5]int{50, 75, 90, 100, 25}
	values := div(arr)
	for i := 0; i < 5; i++ {
		fmt.Println(values[i])
	}
}

func div(numbers [5]int) [5]float64 {
	var arr2 [5]float64
	for i := 0; i < 5; i++ {
		arr2[i] = float64(numbers[i]) / 3
	}
	return arr2
}

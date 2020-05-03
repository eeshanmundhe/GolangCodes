package main

import "fmt"

func avg(num ...float64) float64 {
	var sum float64 = 0
	for _, n := range num {
		fmt.Println(n)
		sum += n
	}
	return sum / float64(len(num))
}

func main() {
	fmt.Println(avg(75, 125))
	fmt.Println(avg(23.3, 36.7))
}

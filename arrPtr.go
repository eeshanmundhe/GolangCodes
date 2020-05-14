package main

import "fmt"

func updatearray(funarr [5]int) {
	funarr[4] = 75
}

func main() {
	arr := [5]int{5, 10, 15, 25, 50}
	updatearray(arr)
	fmt.Println(arr)
}

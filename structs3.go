package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func assign_values(sub subscriber) subscriber {
	sub.name = "Eeshan"
	sub.rate = 7500
	sub.active = false
	return sub
}
func main() {
	var subscriber1 subscriber
	subscriber1 = assign_values(subscriber1)
	fmt.Println(subscriber1)
}

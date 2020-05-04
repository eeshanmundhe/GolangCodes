package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	var subscriber1 subscriber
	var subscriber2 subscriber
	subscriber1.name = "Eeshan"
	subscriber2.name = "Ash"
	subscriber1.rate = 7500
	subscriber2.rate = 8000
	subscriber1.active = false
	subscriber2.active = true
	fmt.Println(subscriber1, subscriber2)
}

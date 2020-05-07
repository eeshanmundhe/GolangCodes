package main

import "fmt"

type button string

func (b *button) switchs() {
	if *b == "turned on" {
		*b = "turned off"
	} else {
		*b = "turned on"
	}
	fmt.Println(*b)
}

type switching interface {
	switchs()
}

func main() {
	b := button("turned off")
	var t switching = &b
	t.switchs()
	t.switchs()
}

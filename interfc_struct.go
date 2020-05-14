package main

import "fmt"

type shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	side int
}

func (s *Square) Area() float64 {
	return (float64)(s.side * s.side)
}

func (s *Square) Perimeter() float64 {
	return (float64)(4 * s.side)
}

func main() {
	var s shape
	s = &Square{side: 4}

	fmt.Println(s.Area())

	fmt.Println(s.Perimeter())
}

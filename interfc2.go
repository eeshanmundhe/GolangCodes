package main

import (
	"fmt"
)

type voices string

func (w voices) crassCacophony() {
	fmt.Println("ffffffs")
}

type horn string

func (h horn) crassCacophony() {
	fmt.Println("Boom")
	fmt.Println(h)
}

type noiseMaker interface {
	crassCacophony()
}

func main() {
	var bird noiseMaker
	bird = voices("robin")
	bird.crassCacophony()
	bird = horn("bluejay")
	bird.crassCacophony()
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, My name is Eeshan."
	fmt.Println(str)
	switching := strings.NewReplacer(",", "!")
	newstr := switching.Replace(str)
	fmt.Println(newstr)
}

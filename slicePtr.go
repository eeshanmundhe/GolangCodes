package main

import (
	"fmt"
)

func main() {
	var s = []string{"A", "B", "C"}
	modifySlice(s)
	fmt.Println(s[0])
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "R"
}

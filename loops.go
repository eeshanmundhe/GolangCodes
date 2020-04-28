package main

import "fmt"

func main() {
	loopfunction("You just made the list", 5)
}

func loopfunction(str string, counter int) {
	for i := 1; i <= counter; i++ {
		fmt.Println(str)
	}
}

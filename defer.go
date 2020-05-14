package main

import "fmt"

func main() {
	var number = 50
	f1()
	defer f2(number)
	number = 100
	fmt.Println("End of main is ", number)
}

func f1() {
	fmt.Println("Hello Friendss! Chai peelo!")
	fmt.Println("End of f1")
}

func f2(number int) {
	fmt.Println("End of f2 is ", number)
}

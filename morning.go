package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	var timeNow int = t.Hour()
	fmt.Println(timeNow)
	if t.Hour() < 12 {
		fmt.Println("Good Morning")
	} else if t.Hour() >= 12 && t.Hour() < 17 {
		fmt.Println("Good afternoon!")
	} else if t.Hour() >= 17 && t.Hour() <= 20 {
		fmt.Println("Good Evening!")
	} else {
		fmt.Println("Good Night!")
	}
}

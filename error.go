package main

import (
	"fmt"
)

type myError struct {
	s string
}

func (err *myError) Error() string {
	return err.s
}

func main() {
	quotient, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error occured: ", err)
		return
	}
	fmt.Println("result is : ", quotient)

}

func divide(num, den int) (int, error) {
	if den == 0 {
		return 0, &myError{s: "can't divide by 0"}
	}
	return num / den, nil
}

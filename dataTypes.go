package main

import "fmt"

func main() {
	var num bool = false
	modifyVal(num)
	//Print num
	fmt.Println(num)
	modifyPtr(&num)
	//print num
	fmt.Println(num)

}
func modifyVal(n bool) {
	n = true
}

func modifyPtr(n *bool) {
	*n = false
}

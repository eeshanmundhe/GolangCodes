package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Printing data in the file")
	file, e := os.Open("values.txt")
	if e != nil {
		log.Fatal(e)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	e = file.Close()
	if e != nil {
		log.Fatal(e)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}

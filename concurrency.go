package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(-1)
	wg.Add(2)
	go delayedIteration1()
	go delayedIteration2()
	wg.Wait()
}

func delayedIteration1() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Inside 1:- Second = ", i)
	}
	wg.Done()
}

func delayedIteration2() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Inside 2:- Second = ", i)
	}
	wg.Done()
}

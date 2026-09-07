package main

import (
	"fmt"
	"time"
)

func runForTwoSec() {
	time.Sleep(2 * time.Second)
	fmt.Println("This is from runForTwoSec funtion")
}

func runForFourSec() {
	time.Sleep(4 * time.Second)
	fmt.Println("This is from runForFourSec funtion")
}

func runForSixSec() {
	time.Sleep(6 * time.Second)
	fmt.Println("This is from runForSixSec funtion")
}

func add(x, y int, ch chan int) {
	time.Sleep(6 * time.Second)
	ch <- x + y
}

func main() {
	ch := make(chan int)
	go add(12, 10, ch)
	ch2 := make(chan int)
	go add(10, -10, ch2)
	go runForTwoSec()
	go runForFourSec()
	go runForSixSec()
	time.Sleep(7 * time.Second)
	fmt.Println("Done!")

	for i := 0; i < 2; i++ {
		select {
		case z := <-ch:
			fmt.Println(z)
		case y := <-ch2:
			fmt.Println(y)
		}
	}
}

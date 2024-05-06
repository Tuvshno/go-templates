package main

import (
	"fmt"
	"time"
)

func gen(ch chan int) {
	sum := 0
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		sum = sum + 1
	}
	ch <- sum
}

func main() {
	ch := make(chan int)

	go gen(ch)

	fmt.Println("main is waiting")
	result := <-ch

	fmt.Println(result)
}

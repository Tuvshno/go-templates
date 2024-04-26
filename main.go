package main

import (
	"fmt"
	"strconv"
	"time"
)

func sendNumbers(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- i
	}

	fmt.Println("No More Numbers")

}

func sendMsgs(ch chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- strconv.Itoa(i)
	}

	fmt.Println("No More Msgs")

}

func main() {
	numbers := make(chan int)
	msgs := make(chan string)

	go sendNumbers(numbers)
	go sendMsgs(msgs)

	for i := 0; i < 10; i++ {
		select {
		case num := <-numbers:
			fmt.Printf("number %d\n", num)
		case msg := <-msgs:
			fmt.Printf("msg %s\n", msg)
		}
	}

}

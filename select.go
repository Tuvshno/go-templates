package main

import (
	"fmt"
	"strconv"
)

func sendNum(ch chan int) {
	// time.Sleep(time.Second * 5)

	for i := 0; i < 5; i++ {
		ch <- i
	}
	fmt.Println("No more numbers")
	close(ch)
}

func sendMsgs(ch chan string) {
	// time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		ch <- strconv.Itoa(i)
	}

	fmt.Println("No more msg")
	close(ch)
}

func main() {
	numbers := make(chan int)
	msgs := make(chan string)

	go sendNum(numbers)
	go sendMsgs(msgs)

	numClosed, msgClosed := false, false

	for !numClosed || !msgClosed {
		select {
		case n, ok := <-numbers:
			if ok {
				fmt.Printf("%d", n)
			} else {
				numClosed = true
			}
		case s, ok := <-msgs:
			if ok {
				fmt.Printf("%s", s)
			} else {
				msgClosed = true
			}
		default:
			// fmt.Println("Nothing recieved yet")
		}
	}

}

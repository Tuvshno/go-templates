package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case i := <-ch:
		fmt.Println(i)
	default:
		fmt.Println("Nothing recieved")
	}

	select {
	case ch <- 42:
		fmt.Println("Sent 42")
	default:
		fmt.Println("No receiver to send to")
	}

}

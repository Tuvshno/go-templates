package main

import (
	"fmt"
	"time"
)

func quicker(quickTimer *time.Ticker) {
	for {
		select {
		case x := <-quickTimer.C:
			fmt.Println("quick", x)
		}
	}
}

func slower(slowTimer *time.Timer) {
	for {
		select {
		case x := <-slowTimer.C:
			fmt.Println("slower", x)
		}
	}
}

func main() {
	quick := time.NewTimer(time.Second)
	slow := time.NewTimer(time.Second * 5)
	stopper := time.NewTimer(time.Second * 4)

	go quicker(quick)
	go slower(slow)

	select {
	case <-stopper.C:
		quick.Stop()
		stopped := slow.Stop()
		fmt.Println(stopped)
	}

}

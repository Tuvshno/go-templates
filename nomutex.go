package main

import (
	"fmt"
	"time"
)

func writer(x map[int]int, factor int) {
	i := 1
	for {
		time.Sleep(time.Second)
		x[i] = x[i-1] * factor
		i++
	}
}

func reader(x map[int]int) {
	for {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(x)
	}
}

func main() {
	x := make(map[int]int)
	x[0] = 1

	go writer(x, 2)
	go reader(x)

	time.Sleep(time.Millisecond * 300)
	go writer(x, 3)

	time.Sleep(time.Second * 4)
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var first int

func setter(id int, ch chan bool, once *sync.Once) {
	t := rand.Uint32() % 300
	time.Sleep(time.Duration(t) * time.Millisecond)
	once.Do(func() {
		first = id
	})

	ch <- true

	fmt.Println(id, "Done")
}

func main() {
	// rand.Seed(time.Now().UnixNano())

	var once sync.Once

	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go setter(i, ch, &once)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}

	fmt.Println("The first was", first)
}

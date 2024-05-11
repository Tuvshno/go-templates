package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	number, err := xml.Marshal(42)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(number))

	msg, _ := xml.Marshal("This is a msg!!!!")
	fmt.Println(string(msg))

}

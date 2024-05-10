package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	msg := "Save the world with Go!!!"
	filePath := "./msg"

	err := ioutil.WriteFile(filePath, []byte(msg), 0644)
	if err != nil {
		panic(err)
	}

	read, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", read)
}

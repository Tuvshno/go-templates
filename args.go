package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	argsWithProg := os.Args

	if len(argsWithProg) <= 1 {
		fmt.Println("Error: Add two arguments")
		os.Exit(2)
	}

	//
	numA, err := strconv.Atoi(argsWithProg[1])

	// If there is an error
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	numB, err := strconv.Atoi(argsWithProg[2])

	// If there is an error
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	result := numA + numB
	fmt.Printf("%d + %d = %d\n", numA, numB, result)

}

package main

import (
	"encoding/csv"
	"os"
)

func main() {
	out := [][]string{
		{"userid", "score", "pass"},
		{"Gopher", "100", "admin"},
	}

	writer := csv.NewWriter(os.Stdout)

	for _, i := range out {
		err := writer.Write(i)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
} 

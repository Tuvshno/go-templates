package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type User struct {
	User     string `json:"username,omitempty"`
	Score    int    `json:"score,omitempty"`
	password string `json:"password, omitempty"`
}

func main() {
	userA := User{"Gopher", 1000, "admin"}
	userB := User{"JJ", 600, "1234"}
	userC := User{User: "Gopher", password: "admin"}

	db := []User{userA, userB, userC}
	dbJson, err := json.Marshal(&db)
	if err != nil {
		panic(err)
	}

	var recovered []User
	err = json.Unmarshal(dbJson, &recovered)
	if err != nil {
		panic(err)
	}
	fmt.Println(recovered)

	var indented bytes.Buffer
	err = json.Indent(&indented, dbJson, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(indented.String())

}

package main

import (
	"os"
	"text/template"
)

type User struct {
	Name   string
	UserId string
	Email  string
}

const Msg = `Dear {{.Name}},
Your id is {{.UserId}} 
and email {{.Email}}
`

func main() {
	u := User{"somename", "useridsome", "some@gmail.com"}
	t := template.Must(template.New("msg").Parse(Msg))
	err := t.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}
}

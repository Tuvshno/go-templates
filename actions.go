package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Score uint32
}

const Msg = `
	{{.Name}} your score is  {{.Score}}
	your level is :
	{{if le .Score 50}} Amateur
	{{else if le .Score 80}} Professional
	{{else}} Expert
	{{end}}
	`

const Musk = `
 musketeers:
 {{range .}} {{print .}}{{end}}
`

func main() {
	u1 := User{"John", 30}
	u2 := User{"Monn", 80}

	t := template.Must(template.New("score").Parse(Msg))
	err := t.Execute(os.Stdout, u1)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, u2)
	if err != nil {
		panic(err)
	}

	musketeers := []string{"m1", "m2", "m3"}

	t = template.Must(template.New("musk").Parse(Musk))
	err = t.Execute(os.Stdout, musketeers)
	if err != nil {
		panic(err)
	}
}

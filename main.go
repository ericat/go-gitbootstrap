package main

import (
	"os"
	"text/template"

	"log"
)

type Settings struct {
	Username string
	Email    string
	Editor   string
}

func main() {
	s := Settings{}
	err := ReadUserInput(s)

	const gitconfig = `
[user]
  name = {{.Username -}}
  email = {{.Email -}}
[core]
  editor = {{.Editor -}}
  whitespace = fix,trailing-space,cr-at-eol
[push]
  default = current
[help]
  autocorrect = 1
`
	t, err := template.New("settings").Parse(gitconfig)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(".gitconfig")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, s)

	if err != nil {
		log.Fatal(err)
	}
}

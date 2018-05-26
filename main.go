package main

import (
  "os"
  "text/template"

  "log"
  "github.com/ttacon/chalk"
  "gopkg.in/AlecAivazis/survey.v1"
)

var qs = []*survey.Question{
  {
    Name:     "username",
    Prompt:   &survey.Input{Message: "What is your github username?"},
    Validate: survey.Required,
    Transform: survey.Title,
  },
  {
    Name: "email",
    Prompt:   &survey.Input{Message: "What is your github email?"},
  },
  {
    Name: "editor",
    Prompt: &survey.Select{
      Message: "What is your favourite editor?",
      Options: []string{"Vim", "VS Code", "Atom", "Sublime"},
      Default: "Vim",
    },
  },
}

func main() {
  answers := struct {
    Username          string
    Email           string
    Editor string `survey:"editor"`
  }{}

  err := survey.Ask(qs, &answers)
  if err != nil {
    log.Fatal(err.Error())
    return
  }
  const gitconfig = `
  [user]
  name = {{.Username}}
  email = {{.Email}}
  [core]
  editor = {{.Editor}}
  whitespace = fix,trailing-space,cr-at-eol
  [push]
  default = current
  [help]
  autocorrect = 1
  `

  t, err := template.New("answers").Parse(gitconfig)
  if err != nil {
    panic(err)
  }

  f, err := os.Create(".gitconfig")
  if err != nil {
    log.Println("create file: ", err)
    return
  }

  err = t.Execute(f, answers)

  if err != nil {
    log.Fatal(err)
  }
}

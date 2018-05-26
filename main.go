package main

import (
	"os"
  "io/ioutil"
	"text/template"
	"log"

	"gopkg.in/AlecAivazis/survey.v1"
)

var qs = []*survey.Question{
	{
		Name:      "username",
		Prompt:    &survey.Input{Message: "What is your github username?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:   "email",
		Prompt: &survey.Input{Message: "What is your github email?"},
	},
	{
		Name: "editor",
		Prompt: &survey.Select{
			Message: "What is your favourite editor?",
			Options: []string{"Vim", "Emacs", "VS Code", "Atom", "Sublime"},
			Default: "Vim",
		},
	},
}

func main() {
	answers := struct {
		Username string
		Email    string
		Editor   string `survey:"editor"`
	}{}

  editors := map[string]string{
    "Vim": "vim",
    "Emacs": "emacs",
    "VS Code": "code --wait",
    "Atom": "atom --wait",
    "Sublime": "subl -n -w",
  }

	err := survey.Ask(qs, &answers)
  answers.Editor = editors[answers.Editor]
	if err != nil {
		log.Fatal(err.Error())
		return
	}

  file, err := ioutil.ReadFile("template.txt")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("answers").Parse(string(file))
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(".gitconfig")
	if err != nil {
	}

	err = t.Execute(f, answers)

	if err != nil {
		log.Fatal(err)
	}
}

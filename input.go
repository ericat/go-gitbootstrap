package main

import (
	"bufio"
	"fmt"
)

func ReadUserInput(s Settings) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What is your Github username?")
	s.Username, _ = reader.ReadString('\n')
	if len(s.Username) <= 1 {
		return errors.New("Your username is required!")
	}

	fmt.Println("What is your Github email?")
	s.Email, _ = reader.ReadString('\n')
	if len(s.Email) <= 1 {
		return errors.New("Your email is required!")
	}

	fmt.Println("What is your favourite editor?")
	s.Editor, _ = reader.ReadString('\n')

	return nil
}

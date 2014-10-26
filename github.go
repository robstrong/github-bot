package main

import (
	"bufio"
	"code.google.com/p/gopass"
	"fmt"
	"os"
)

func PromptForCredentials() (user string, pass string, err error) {
	user, err = getUserName()
	if err != nil {
		return
	}
	pass, err = gopass.GetPass("Enter GitHub Password (this will not be stored):")
	return
}

func getUserName() (username string, err error) {
	fmt.Println("Enter your GitHub username:")
	bio := bufio.NewReader(os.Stdin)
	user, _, err := bio.ReadLine()
	username = string(user[:])
	return
}

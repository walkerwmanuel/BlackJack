package userInput

import (
	"bufio"
	"fmt"
	"os"
)

// GetUIForUsername - Ask for user to inpt username and stores that input to a string
func GetUIForUsername() string {
	fmt.Print("Enter username: ")
	input := bufio.NewReader(os.Stdin)
	username, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}
	return username
}

// GetUIForPassword - Ask for user to inpt password and stores that input to a string
func GetUIForPassword() string {
	fmt.Print("Enter password: ")
	input2 := bufio.NewReader(os.Stdin)
	password, err := input2.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}
	return password
}

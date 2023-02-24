package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"walkerwmanuel/blackjack/create"
	"walkerwmanuel/blackjack/data"
)

func main() {
	//Calls to connect to database
	err := data.ConnectDatabase()
	if err != nil {
		fmt.Println("Error connecting to database")
		os.Exit(1)
	}

	//Ask for user input on username
	fmt.Print("Enter username: ")
	input := bufio.NewReader(os.Stdin)
	username, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	//Ask for user input on password
	fmt.Print("Enter password: ")
	input2 := bufio.NewReader(os.Stdin)
	password, err := input2.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	x := create.CreatePlayer(username, password)

	fmt.Printf("Created Player: %v ", x)

	y := create.CreateGame(1)

	z := create.AddPlayerToGame(x, y)

	z1, err := json.Marshal(z)
	if err != nil {
		fmt.Println("Error trynna do that json")
	}

	z2 := string(z1[:])

	fmt.Printf("Game with player added to it: %v", z2)

}

// func main2() {
// 	fmt.Print("Enter text: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	// ReadString will block until the delimiter is entered
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("An error occured while reading input. Please try again", err)
// 		return
// 	}

// 	// remove the delimeter from the string
// 	input = strings.TrimSuffix(input, "\n")
// 	fmt.Println(input)
// }

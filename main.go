package main

import (
	"encoding/json"
	"fmt"
	"walkerwmanuel/blackjack/create"
	"walkerwmanuel/blackjack/data"
	"walkerwmanuel/blackjack/userInput"
)

func main() {
	// Calls Connect to database
	err := data.ConnectDatabase()
	if err != nil {
		fmt.Printf("%v \n", err)
		return
	}
	fmt.Println("Connected to the Database!")

	//Ask for user input for player struct vales
	username := userInput.GetUIForUsername()
	password := userInput.GetUIForPassword()

	x := create.CreatePlayer(username, password)

	fmt.Printf("Created Player: %v ", x)

	//
	//
	//
	//
	//Game logic below
	//
	//
	//
	//

	y := create.CreateGame(1)

	z := create.AddPlayerToGame(x, y)

	z1, err := json.Marshal(z)
	if err != nil {
		fmt.Println("Error trynna do that json")
	}

	z2 := string(z1[:])

	fmt.Printf("Game with player added to it: %v", z2)

}

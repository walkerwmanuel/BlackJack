package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"walkerwmanuel/blackjack/data"
	"walkerwmanuel/blackjack/logic"
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

	err = data.CreateTableGames()
	if err != nil {
		fmt.Println("error")
	}

	err = data.CreateTablePlayers()
	if err != nil {
		fmt.Println("error")
	}
	//Ask for user input for player struct vales
	username := userInput.GetUIForUsername()
	password := userInput.GetUIForPassword()

	x := logic.CreatePlayer(username, password)

	fmt.Printf("Created Player: %v ", x)

	gameId := userInput.GetUIForGameId()

	//Converts input from byte to int
	gameIdInt, err := strconv.Atoi(string(gameId))
	if err != nil {
		fmt.Println(err)
	}

	//Creates game based on input
	y := logic.CreateGame(gameIdInt)

	z := logic.AddPlayerToGame(x, y)

	z1, err := json.Marshal(z)
	if err != nil {
		fmt.Println(err)
	}

	z2 := string(z1[:])

	fmt.Printf("Game with player added to it: %v", z2)

	_, err = data.InsertGameToDB(z)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Your game has been added to the database!")
}

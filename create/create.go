package create

import (
	"walkerwmanuel/blackjack/types"
)

// CreatePlayer - takes in a username and password and put it into a pointer of a player struct
func CreatePlayer(u, p string) *types.Player {

	newPlayer := types.Player{}

	newPlayer.Username = u
	newPlayer.Password = p

	return &newPlayer
}

// CreateGame - takes in an id and returns a game with a blank non nil array of players
func CreateGame(i int) *types.Game {

	newGame := types.Game{}

	newGame.Id = i
	// Make array with 0 memory allocation
	newGame.Players = []*types.Player{}

	return &newGame
}

// AddPlayerToGame - Takes in game and player struct, adds the player to the game struct in its player value
func AddPlayerToGame(newPlayer *types.Player, newGame *types.Game) *types.Game {

	newGameWithPlayer := types.Game{}

	newGameWithPlayer.Id = newGame.Id
	newGameWithPlayer.Players = append(newGameWithPlayer.Players, newPlayer)

	return &newGameWithPlayer
}

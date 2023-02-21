package create

import (
	"walkerwmanuel/blackjack/types"
)

func CreatePlayer(u, p string) *types.Player {

	newPlayer := types.Player{}

	newPlayer.Username = u
	newPlayer.Password = p

	return &newPlayer
}

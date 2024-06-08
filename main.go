package main

import (
	"MagicalArena/arena"
)

func main() {
	playerA := arena.CreatePlayer("Player A")
	playerB := arena.CreatePlayer("Player B")

	game := arena.NewGame(playerA, playerB)
	game.Start()
}



package main

import (
	"dnd-web-interface/internal/domain"
	"fmt"
)

func main() {
	player := domain.NewPlayer(
		"Adam",
		domain.Adventurer,
		domain.NewCharacter(
			"Bob",
			1,
			10,
			10,
			10,
			10,
			10,
			10,
		),
	)
	fmt.Println(player)
	fmt.Println(player.Character)
}

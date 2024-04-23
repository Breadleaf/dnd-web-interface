package main

import (
	"dnd-web-interface/internal/domain"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
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

	// I am truely sorry for this
	playerString := func(p *domain.Player) string {
		return fmt.Sprintf(
			"Player Name: %s\n" +
			"Player Type: %d\n" +
			"Character:\n" +
			"- Name: %s\n" +
			"- Level: %d\n" +
			"- Strength: %d\n" +
			"- Dexterity: %d\n" +
			"- Constitution: %d\n" +
			"- Intelligence: %d\n" +
			"- Wisdom: %d\n" +
			"- Charisma: %d\n",
			p.PlayerName,
			p.PlayerType,
			p.Character.Name,
			p.Character.Level,
			p.Character.Strength,
			p.Character.Dexterity,
			p.Character.Constitution,
			p.Character.Intelligence,
			p.Character.Wisdom,
			p.Character.Charisma,
		)
	}

	router := mux.NewRouter()

	router.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		// This serves as a simple hello world
		fmt.Fprintf(w, playerString(player))
	})

	http.ListenAndServe(":8080", router)
}

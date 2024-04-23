package domain

type PlayerType int

const (
	DungeonMaster PlayerType = iota
	Adventurer
)

type Player struct {
	PlayerName string
	PlayerType PlayerType
	Character *Character
}

func NewPlayer(playerName string, playerType PlayerType, character *Character) *Player {
	return &Player {
		PlayerName: playerName,
		PlayerType: playerType,
		Character: character,
	}
}

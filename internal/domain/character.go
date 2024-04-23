package domain

type Character struct {
	Name string
	Level int
	Strength int
	Dexterity int
	Constitution int
	Intelligence int
	Wisdom int
	Charisma int
}

func NewCharacter(name string, level int, strength int, dexterity int, constitution int, intelligence int, wisdom int, charisma int) *Character {
	return &Character{
		Name: name,
		Level: level,
		Strength: strength,
		Dexterity: dexterity,
		Constitution: constitution,
		Intelligence: intelligence,
		Wisdom: wisdom,
		Charisma: charisma,
	}
}

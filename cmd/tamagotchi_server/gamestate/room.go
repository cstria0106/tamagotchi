package gamestate

type Room struct {
	ID        uint32
	Foods     []*Food
	Poos      []*Poo
	Character *Character
}

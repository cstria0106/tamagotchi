package gamestate

type Character struct {
	ID    uint32
	Doing bool
	Point Point
	Full  uint8
}

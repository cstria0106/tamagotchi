package game

type System interface {
	Init(g Game) error
	Update(g Game) error
}

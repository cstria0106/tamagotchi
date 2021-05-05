package server

import (
	"sync"
	"tamagotchi/network/events/buffers/serverbuffer"
	"tamagotchi/server/gamestate"
	"time"
)

type Game struct {
	server *Server
	room   *gamestate.Room
	lastId uint32
}

func (g *Game) getNextID() uint32 {
	g.lastId++
	return g.lastId
}

func CreateGame(server *Server) *Game {
	return &Game{
		server: server,
		room: &gamestate.Room{
			Foods: []*gamestate.Food{},
			Poos:  []*gamestate.Poo{},
			Character: &gamestate.Character{
				ID:    0,
				Doing: false,
				Point: gamestate.Point{},
				Full:  80,
			},
		},
	}
}

func (g *Game) Start() {
	for {
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			g.onTick()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 60)
			wg.Done()
		}()

		wg.Wait()
	}
}

func (g *Game) onTick() {
	if !g.room.Character.Doing {
		if g.room.Character.Full > 80 {
			go g.poop()
		} else if len(g.room.Foods) > 0 {
			go g.eatFood(g.room.Foods[0])
		}
	}

}

func (g *Game) poop() {
	g.room.Character.Doing = true

	time.Sleep(time.Second * 1)
	g.room.Character.Full -= 10
	g.addPoo(g.room.Character.Point)

	g.room.Character.Doing = false
}

func (g *Game) eatFood(food *gamestate.Food) {
	g.room.Character.Doing = true

	g.room.Character.Point = food.Point
	g.server.sendToAll(serverbuffer.CharacterMoveBuffer(food.Point.X, food.Point.Y))

	time.Sleep(time.Second)

	for i, f := range g.room.Foods {
		if f == food {
			g.room.Foods = append(g.room.Foods[:i], g.room.Foods[i+1:]...)
			break
		}
	}

	g.room.Character.Full += 10

	g.server.sendToAll(serverbuffer.CharacterEatBuffer(food.ID))
	g.room.Character.Doing = false
}

func (g *Game) addPoo(point gamestate.Point) {
	poo := &gamestate.Poo{
		ID:    g.getNextID(),
		Point: point,
	}
	g.room.Poos = append(g.room.Poos, poo)

	g.server.sendToAll(serverbuffer.AddPooBuffer(poo.ID, point.X, point.Y))
}

func (g *Game) addFood(point gamestate.Point) {
	food := &gamestate.Food{
		ID:    g.getNextID(),
		Point: point,
	}
	g.room.Foods = append(g.room.Foods, food)

	g.server.sendToAll(serverbuffer.AddFoodBuffer(food.ID, point.X, point.Y))
}

func (g *Game) clean(id uint32) {
	removed := false
	for i, poo := range g.room.Poos {
		if poo.ID == id {
			g.room.Poos = append(g.room.Poos[:i], g.room.Poos[i+1:]...)
			removed = true
			break
		}
	}

	if !removed {
		return
	}

	g.server.sendToAll(serverbuffer.CleanPooBuffer(id))
}

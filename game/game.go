package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
	"image/color"
	"log"
	"os"
	"tamagotchi/game/client"
	"tamagotchi/game/drawable"
	"tamagotchi/network/events"
	"tamagotchi/network/events/buffers/clientbuffer"
	"tamagotchi/resources/images"
	"tamagotchi/util"
	"time"
)

const (
	screenWidth  = 200
	screenHeight = 150
	screenScale  = 4
)

type game struct {
	client        client.Client
	drawables     []drawable.Drawable
	drawableCount int

	cleanButton, feedButton *drawable.Button
	poos                    []*drawable.Poo
	foods                   []*drawable.Food
}

func (g *game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	g.cleanButton.CheckClick()
	g.feedButton.CheckClick()

	if g.cleanButton.JustDown {
		if len(g.poos) > 0 {
			g.client.Send(clientbuffer.CleanBuffer(g.poos[0].ID))
		}
	}

	if g.feedButton.JustDown {
		g.client.Send(clientbuffer.FeedBuffer(
			uint16(rand.Uint32()%200),
			uint16(rand.Uint32()%150),
		))
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for _, d := range g.drawables[:g.drawableCount] {
		d.DrawOn(screen)
	}
}

func (g *game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *game) addDrawable(d drawable.Drawable) {
	g.drawables[g.drawableCount] = d
	g.drawableCount++
}

func (g *game) removeDrawable(d drawable.Drawable) {
	for i, element := range g.drawables {
		if element == d {
			g.drawables = append(g.drawables[:i], g.drawables[i+1:]...)
			g.drawableCount--
			return
		}
	}
}

func StartGame() {
	rand.Seed(uint64(time.Now().Unix()))

	c := client.Connect()

	character := drawable.NewCharacter(screenWidth/2, screenHeight/2)
	cleanButton := drawable.NewButtonFromImage(images.CleanButton, screenWidth/2+1, screenHeight-15, 1, 1)
	feedButton := drawable.NewButtonFromImage(images.FeedButton, screenWidth/2-29, screenHeight-15, 1, 1)

	g := &game{
		client:        c,
		drawables:     make([]drawable.Drawable, 1024),
		drawableCount: 0,
		cleanButton:   cleanButton,
		feedButton:    feedButton,
		poos:          []*drawable.Poo{},
		foods:         []*drawable.Food{},
	}

	c.AddListener(events.CleanPoo, func(buffer []byte) {
		id := util.DecodeU32(buffer)
		for i, poo := range g.poos {
			if poo.ID == id {
				g.removeDrawable(poo)
				g.poos = append(g.poos[:i], g.poos[i+1:]...)
				return
			}
		}
	})

	c.AddListener(events.CharacterEat, func(buffer []byte) {
		id := util.DecodeU32(buffer)
		for i, food := range g.foods {
			if food.ID == id {
				g.removeDrawable(food)
				g.foods = append(g.foods[:i], g.foods[i+1:]...)
				return
			}
		}
	})

	c.AddListener(events.AddFood, func(buffer []byte) {
		food := drawable.NewFood(
			util.DecodeU32(buffer[0:4]),
			util.DecodeI16(buffer[4:6]),
			util.DecodeI16(buffer[6:8]),
		)
		g.addDrawable(food)
		g.foods = append(g.foods, food)
	})

	c.AddListener(events.AddPoo, func(buffer []byte) {
		poo := drawable.NewPoo(
			util.DecodeU32(buffer[0:4]),
			util.DecodeI16(buffer[4:6]),
			util.DecodeI16(buffer[6:8]),
		)
		g.addDrawable(poo)
		g.poos = append(g.poos, poo)
	})

	c.AddListener(events.CharacterMove, func(buffer []byte) {
		x := util.DecodeI16(buffer[0:2])
		y := util.DecodeI16(buffer[2:4])
		character.MoveTo(x, y)
	})

	g.addDrawable(character)
	g.addDrawable(cleanButton)
	g.addDrawable(feedButton)

	ebiten.SetWindowSize(screenWidth*screenScale, screenHeight*screenScale)
	//ebiten.SetScreenTransparent(true)
	//ebiten.SetWindowDecorated(false)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

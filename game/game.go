package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"os"
	"tamagotchi/network"
	"tamagotchi/resources/images"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	client        Client
	drawables     []Drawable
	drawableCount int

	feedButton, cleanButton *Button
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	g.cleanButton.CheckClick()
	g.feedButton.CheckClick()

	if g.cleanButton.JustDown {
		g.client.Send(network.BuildPayload(network.ActionClean, nil))
	}

	if g.feedButton.JustDown {
		// TODO Queue 를 분리하여 읽도록 서버를 수정해야 한다.
		for {
			go g.client.Send(network.BuildPayload(network.ActionFeed, []byte("FEED OK!")))
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, drawable := range g.drawables[:g.drawableCount] {
		drawable.DrawOn(screen)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) addDrawable(d Drawable) {
	g.drawables[g.drawableCount] = d
	g.drawableCount++
}

func StartGame() {
	client := Connect()

	character := NewCharacter(screenWidth/2, screenHeight/2)
	feedButton := NewButtonFromImage(images.FeedButton, screenWidth/2-64-128, screenHeight-128, 1, 1)
	cleanButton := NewButtonFromImage(images.CleanButton, screenWidth/2-64+128, screenHeight-128, 1, 1)

	game := &Game{
		client:        client,
		drawables:     make([]Drawable, 1024),
		drawableCount: 0,
		feedButton:    feedButton,
		cleanButton:   cleanButton,
	}

	game.addDrawable(character)
	game.addDrawable(feedButton)
	game.addDrawable(cleanButton)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

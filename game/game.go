package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"os"
	"tamagotchi/resources/images"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
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

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, drawable := range g.drawables[:g.drawableCount] {
		drawable.DrawOn(screen)
	}

	if g.cleanButton.JustDown {
		ebitenutil.DebugPrintAt(screen, "Clean!", 200, 0)
	}

	if g.feedButton.JustDown {
		ebitenutil.DebugPrintAt(screen, "Feed!", 200, 20)
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
	character := NewCharacter(screenWidth/2, screenHeight/2)
	feedButton := NewButtonFromImage(images.FeedButton, screenWidth/2-64-128, screenHeight-128, 1, 1)
	cleanButton := NewButtonFromImage(images.CleanButton, screenWidth/2-64+128, screenHeight-128, 1, 1)

	game := &Game{
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

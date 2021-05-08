package tamagotchi

import (
	"tamagotchi/cmd/tamagotchi/client"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/tamagotchi/entity"
	"tamagotchi/cmd/tamagotchi/tamagotchi/system"
)

const (
	screenWidth  = 100
	screenHeight = 150
	screenScale  = 4
)

func Start(c *client.Client) error {
	g, err := game.New(&game.Options{
		Client: c,
		Screen: &game.ScreenOptions{
			Width:  screenWidth,
			Height: screenHeight,
			Scale:  screenScale,
		},
	})

	if err != nil {
		return err
	}

	background := game.MergeSystems(system.Background(), system.FPSCounter())
	foreground := game.MergeSystems(system.Tween(), system.Mouse(), system.Draw())
	play := game.MergeSystems(system.Character())
	top := game.MergeSystems(system.FPSCounter(), system.Cursor())

	err = g.AddSystem(background, play, foreground, top)
	if err != nil {
		return err
	}

	g.AddEntities(entity.NewCharacter(screenWidth/2, screenHeight/2))
	g.AddEntities(entity.NewCursor())

	return g.Start()
}

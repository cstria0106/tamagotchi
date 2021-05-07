package tamagotchi

import (
	"tamagotchi/cmd/tamagotchi/client"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/tamagotchi/service"
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

	_ = g.AddService(service.NewBasicService(g))

	return g.Start()
}

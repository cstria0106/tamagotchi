package tamagotchi

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/cli"
	"tamagotchi/cmd/tamagotchi/component"
	"tamagotchi/cmd/tamagotchi/engine"
	"tamagotchi/cmd/tamagotchi/entity"
	"tamagotchi/cmd/tamagotchi/system"
	"tamagotchi/internal/game"
)

const (
	screenWidth  = 100
	screenHeight = 150
	//screenScale  = 4
)

func StartTamagotchi(arguments *cli.Arguments) error {
	tamagotchi, err := engine.NewGame(&engine.Options{
		Arguments: arguments,
		Screen: &engine.ScreenOptions{
			Width:  100,
			Height: 150,
			Scale:  4,
		},
	})

	if err != nil {
		return err
	}

	background := engine.MergeSystems(system.Background, system.Cursor, system.Game)
	foreground := engine.MergeSystems(system.Tween, system.Mouse, system.Draw)
	play := engine.MergeSystems(system.Character)
	top := engine.MergeSystems(system.FPSCounter)

	err = tamagotchi.AddSystem(background, play, foreground, top)
	if err != nil {
		return err
	}

	character := entity.NewCharacter(screenWidth/2, screenHeight/2)
	tamagotchi.AddEntities(character)
	tamagotchi.AddEntities(entity.NewCursor())
	tamagotchi.AddEntities(game.NewEntity(game.NewEntityComponents(component.NewNetwork(arguments.Host, arguments.Port))))

	err = tamagotchi.Start()
	if err != nil {
		return err
	}

	return ebiten.RunGame(tamagotchi)
}

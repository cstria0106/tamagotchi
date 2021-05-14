package system

import (
	"fmt"
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/component"
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/engine"
	"github.com/cstria0106/tamagotchi/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Game = engine.NewSystem(&engine.SystemOptions{
	Init: func(g game.Game) error {
		g.WithComponents([]game.ComponentID{component.NETWORK}, func(components []*game.Component) {
			network := components[0].Data.(*component.Network)
			network.ConnectAsync(func() {
				network.PingAsync(nil)
			})
		})

		return nil
	},
	Draw: func(g game.Game, screen *ebiten.Image) {
		g.WithComponents([]game.ComponentID{component.NETWORK}, func(components []*game.Component) {
			network := components[0].Data.(*component.Network)

			if network.Connected {
				s := "connected"

				if network.RemoteVersion != nil {
					s += fmt.Sprintf("\n%s", network.RemoteVersion)
				}

				ebitenutil.DebugPrint(screen, s)
			} else if network.Error != nil {
				ebitenutil.DebugPrint(screen, "error")
			}
		})
	},
})

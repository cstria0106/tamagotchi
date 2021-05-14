package main

import (
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/cli"
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/resources/images"
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/tamagotchi"
	"github.com/sqweek/dialog"
	"os"
)

func preloadImages() error {
	return images.Load()
}

func handleError(err error) {
	if err != nil {
		dialog.Message("could not start engine: %v", err)
		os.Exit(1)
	}
}

func main() {
	arguments := cli.GetArguments()

	handleError(preloadImages())
	handleError(tamagotchi.StartTamagotchi(arguments))
}

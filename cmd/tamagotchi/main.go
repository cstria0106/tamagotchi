package main

import (
	"github.com/sqweek/dialog"
	"os"
	"tamagotchi/cmd/tamagotchi/cli"
	"tamagotchi/cmd/tamagotchi/resources/images"
	"tamagotchi/cmd/tamagotchi/tamagotchi"
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

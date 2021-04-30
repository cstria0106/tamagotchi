package main

import (
	_ "image/png"
	"log"
	"os"
	"tamagotchi/game"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-s" {
		log.Println("Start server...")

		return
	}

	game.StartGame()
}

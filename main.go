package main

import (
	_ "image/png"
	"log"
	"os"
	"tamagotchi/game"
	"tamagotchi/server"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-s" {
		log.Println("Start server...")
		server.StartServer()
		return
	}

	game.StartGame()
}

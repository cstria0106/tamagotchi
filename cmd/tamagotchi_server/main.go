package main

import (
	"tamagotchi/cmd/tamagotchi/cli"
)

func main() {
	arguments := cli.GetArguments()
	StartServer(arguments.Port)
}

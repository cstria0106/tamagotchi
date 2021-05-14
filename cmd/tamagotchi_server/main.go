package main

import (
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/cli"
)

func main() {
	arguments := cli.GetArguments()
	StartServer(arguments.Port)
}

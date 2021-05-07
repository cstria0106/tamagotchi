package main

import (
	"flag"
	"fmt"
	"github.com/sqweek/dialog"
	"tamagotchi/cmd/tamagotchi/client"
	"tamagotchi/cmd/tamagotchi/resources/images"
	"tamagotchi/cmd/tamagotchi/tamagotchi"
	"tamagotchi/internal/data/version"
	"tamagotchi/internal/util/versionutil"
)

func start(host string, port uint64) error {
	var c *client.Client
	var hostVersion *version.Version
	var err error

	if c, hostVersion, err = client.Connect(host, uint16(port)); err != nil {
		return err
	}

	localVersion, err := versionutil.GetVersion()
	if err != nil {
		return fmt.Errorf("could not read local version: %s", err.Error())
	}

	if !hostVersion.Equals(localVersion) {
		return fmt.Errorf("local version(%s) is not matching with remote version(%s)", localVersion, hostVersion)
	}

	return tamagotchi.Start(c)
}

func main() {
	if err := images.Load(); err != nil {
		dialog.Message(fmt.Sprintf("could not start game: %s", err.Error())).Error()
		return
	}

	var host string
	var port uint64
	flag.StringVar(&host, "host", "127.0.0.1", "host to connect")
	flag.Uint64Var(&port, "port", 27775, "port to connect")
	flag.Parse()

	err := start(host, port)

	if err != nil {
		dialog.Message(fmt.Sprintf("could not start game: %s", err.Error())).Error()
		return
	}
}

package main

import (
	"flag"
	"log"
)

func main() {
	var port uint64
	flag.Uint64Var(&port, "port", 27775, "port to connect/host")
	flag.Parse()

	log.Println("starting server")

	StartServer(uint16(port))
}

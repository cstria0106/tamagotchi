package cli

import "flag"

type Arguments struct {
	Port uint16
}

func GetArguments() *Arguments {
	var port uint64
	flag.Uint64Var(&port, "port", 27775, "port to connect/host")
	flag.Parse()

	return &Arguments{
		Port: uint16(port),
	}
}

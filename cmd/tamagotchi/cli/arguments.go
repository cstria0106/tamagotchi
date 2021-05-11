package cli

import "flag"

type Arguments struct {
	Host string
	Port uint16
}

func GetArguments() *Arguments {
	var host string
	var port uint64

	flag.StringVar(&host, "host", "127.0.0.1", "host to connect")
	flag.Uint64Var(&port, "port", 27775, "port to connect")
	flag.Parse()

	return &Arguments{
		Host: host,
		Port: uint16(port),
	}
}

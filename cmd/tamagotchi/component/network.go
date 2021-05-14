package component

import (
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/client"
	"github.com/cstria0106/tamagotchi/internal/data/version"
	"github.com/cstria0106/tamagotchi/internal/game"
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/events/buffers/clientbuffer"
	"sync"
)

type Network struct {
	host          string
	port          uint16
	mutex         sync.Mutex
	RemoteVersion *version.Version
	Client        *client.Client
	Connected     bool
	Error         error
}

func NewNetwork(host string, port uint16) *game.Component {
	return game.NewComponent(
		NETWORK,
		&Network{
			host: host,
			port: port,
		},
	)
}

func (n *Network) Connect() error {
	c, err := client.Connect(n.host, n.port)

	if err != nil {
		return err
	}

	n.Client = c
	n.Connected = true
	return nil
}

func (n *Network) ConnectAsync(cb func()) {
	n.mutex.Lock()
	go func() {
		err := n.Connect()

		if err != nil {
			n.Error = err
			return
		}

		n.Connected = true
		n.mutex.Unlock()

		if cb != nil {
			cb()
		}
	}()
}

func (n *Network) Ping() (*version.Version, error) {
	payload, err := n.Client.Send(clientbuffer.PingBuffer()).Wait(events.Pong)
	if err != nil {
		return nil, err
	}

	return version.FromBuffer(payload), nil
}

func (n *Network) PingAsync(cb func(v *version.Version)) {
	n.mutex.Lock()
	go func() {
		v, err := n.Ping()
		if err != nil {
			n.Error = err
			return
		}

		n.RemoteVersion = v
		n.mutex.Unlock()

		if cb != nil {
			cb(v)
		}
	}()
}

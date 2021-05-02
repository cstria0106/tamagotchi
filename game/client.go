package game

import (
	"encoding/binary"
	"errors"
	"log"
	"net"
	"tamagotchi/network"
)

type Client struct {
	conn      net.Conn
	listeners []*Listener
}

type Listener interface {
	OnPayload(payload network.Payload) bool
}

func (c *Client) Send(payload *network.Payload) {
	buffer := payload.ToBuffer()
	_, err := c.conn.Write(buffer[:6])
	if err != nil {
		log.Println(err)
	}

	_, err = c.conn.Write(buffer[6:])
	if err != nil {
		log.Println(err)
	}
}

func (c *Client) AddListener(listener *Listener) {
	c.listeners = append(c.listeners, listener)
}

func (c *Client) WaitForPayload() (*network.Payload, error) {
	buffer := make([]byte, 6)
	_, err := c.conn.Read(buffer)

	if err != nil {
		log.Println(err)
		return nil, errors.New("could not read from connection")
	}

	action := network.Action(binary.BigEndian.Uint16(buffer[0:2]))
	length := binary.BigEndian.Uint32(buffer[2:6])

	buffer = make([]byte, length)
	_, err = c.conn.Read(buffer)

	if err != nil {
		log.Println(err)
		return nil, errors.New("could not read from connection")
	}

	return &network.Payload{
		Action:     action,
		DataLength: length,
		Data:       buffer,
	}, nil
}

func (c *Client) Close() {
	err := c.conn.Close()
	if err != nil {
		log.Println(err)
	}
}

func Connect() Client {
	conn, err := net.Dial("udp", "127.0.0.1:27775")
	if err != nil {
		log.Fatalln(err)
	}

	client := Client{
		conn:      conn,
		listeners: []*Listener{},
	}

	client.Send(network.BuildPayload(network.ActionPing, nil))

	pong, err := client.WaitForPayload()

	if err != nil {
		log.Fatal(err)
	}

	if pong.Action != network.ActionPong {
		log.Fatalln("ping response is not pong")
	}

	return client
}

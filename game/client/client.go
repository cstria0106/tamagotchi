package client

import (
	"errors"
	"fmt"
	"log"
	"net"
	"tamagotchi/network/events"
	"tamagotchi/network/events/buffers/clientbuffer"
	"tamagotchi/network/header"
	"tamagotchi/util"
	"time"
)

type Listener = struct {
	id   uint32
	emit func(buffer []byte)
}

type Client struct {
	conn           net.Conn
	listeners      map[events.EventType][]*Listener
	lastListenerId map[events.EventType]uint32
}

func (c *Client) Send(buffer []byte) {
	_, err := c.conn.Write(buffer)
	if err != nil {
		log.Println(err)
	}

	log.Println("Sent", events.EventType(util.DecodeU16(buffer[0:2])).String())
}

func (c *Client) AddListener(eventType events.EventType, function func(buffer []byte)) uint32 {
	id, ok := c.lastListenerId[eventType]

	if !ok {
		id = 0
	}

	id++

	if _, ok = c.listeners[eventType]; !ok {
		c.listeners[eventType] = []*Listener{}
	}

	c.lastListenerId[eventType] = id

	c.listeners[eventType] = append(c.listeners[eventType], &Listener{
		id,
		function,
	})

	return id
}
func (c *Client) RemoveListener(eventType events.EventType, id uint32) error {
	if listeners, ok := c.listeners[eventType]; ok {
		i := 0
		listenerCount := len(listeners)
		for i < listenerCount {
			if listeners[i].id == id {
				listeners = append(listeners[:i], listeners[i+1:]...)
				return nil
			}
		}
	} else {
		return errors.New(fmt.Sprintf("listeners for %s is empty", eventType.String()))
	}

	return errors.New(fmt.Sprintf("there is no Listener %d on type %s", id, eventType.String()))
}

func (c *Client) Close() {
	err := c.conn.Close()
	if err != nil {
		log.Println(err)
	}
}

func (c *Client) HandleEvents() {
	buffer := make([]byte, 6)
	length, err := c.conn.Read(buffer[:])

	if length != 6 {
		log.Printf("received receivedHeader length(%d) is not 6\n", length)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}

	receivedHeader := header.FromBuffer(buffer[:])
	log.Println("Got", receivedHeader.Type.String())

	buffer = make([]byte, receivedHeader.Length)
	length, err = c.conn.Read(buffer[:])

	if uint32(length) != receivedHeader.Length {
		log.Printf("recieved length(%d) is mismatching with metadata length(%d)\n", length, receivedHeader.Length)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}

	for _, listener := range c.listeners[receivedHeader.Type] {
		listener.emit(buffer)
	}
}

func Connect() Client {
	conn, err := net.Dial("tcp", "192.168.0.8:27775")
	if err != nil {
		log.Fatalln(err)
	}

	client := Client{
		conn:           conn,
		listeners:      map[events.EventType][]*Listener{},
		lastListenerId: map[events.EventType]uint32{},
	}

	go func() {
		for {
			client.HandleEvents()
		}
	}()

	client.Send(clientbuffer.PingBuffer())

	pongChannel := make(chan error)

	pongListener := func(_ []byte) {
		pongChannel <- nil
	}

	listenerId := client.AddListener(events.Pong, pongListener)

	go func() {
		time.Sleep(3 * time.Second)
		pongChannel <- errors.New("timed out")
	}()

	if err = <-pongChannel; err != nil {
		log.Fatal(err)
	}

	if err = client.RemoveListener(events.Pong, listenerId); err != nil {
		log.Fatal(err)
	}

	return client
}

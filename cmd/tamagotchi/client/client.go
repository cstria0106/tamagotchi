package client

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"tamagotchi/internal/data/version"
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/events/buffers/clientbuffer"
	"tamagotchi/internal/network/header"
	"tamagotchi/internal/util"
	"time"
)

type Listener = struct {
	event events.EventType
	emit  func(buffer []byte)
}

type Client struct {
	conn      net.Conn
	listeners map[events.EventType][]*Listener
	mutex     sync.Mutex
}

type Future struct {
	client *Client
}

func (f *Future) Wait(event events.EventType) ([]byte, error) {
	bufferChannel := make(chan []byte)
	timeoutChannel := make(chan interface{})

	listener := f.client.Listen(event, func(buffer []byte) {
		bufferChannel <- buffer
	})

	defer f.client.RemoveListener(listener)

	go func() {
		time.Sleep(time.Second * 3)
		timeoutChannel <- nil
	}()

	select {
	case <-timeoutChannel:
		return nil, errors.New("timed out")
	case result := <-bufferChannel:
		return result, nil
	}
}

func (c *Client) Send(buffer []byte) *Future {
	_, err := c.conn.Write(buffer[0:6])
	if err != nil {
		log.Println(err)
	}

	if util.DecodeU32(buffer[2:6]) != 0 {
		_, err := c.conn.Write(buffer[6:])
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("Sent", events.EventType(util.DecodeU16(buffer[0:2])).String())

	return &Future{c}
}

func (c *Client) Listen(eventType events.EventType, function func(buffer []byte)) *Listener {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.listeners[eventType]; !ok {
		c.listeners[eventType] = []*Listener{}
	}

	listener := &Listener{
		eventType,
		function,
	}

	c.listeners[eventType] = append(c.listeners[eventType], listener)

	return listener
}
func (c *Client) RemoveListener(listener *Listener) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	event := listener.event

	if listeners, ok := c.listeners[event]; ok {
		for i, l := range listeners {
			if l == listener {
				c.listeners[event] = append(c.listeners[event][:i], c.listeners[event][i+1:]...)
				return true
			}
		}
	}

	return false
}

func (c *Client) Close() {
	err := c.conn.Close()
	if err != nil {
		log.Println(err)
	}
}

func (c *Client) HandleEvents() {
	payload := make([]byte, 6)
	length, err := c.conn.Read(payload[:])

	if length != 6 {
		log.Printf("received receivedHeader length(%d) is not 6\n", length)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}

	receivedHeader := header.FromBuffer(payload[:])
	payload = make([]byte, receivedHeader.Length)

	if receivedHeader.Length > 0 {
		length, err = c.conn.Read(payload[:])

		if uint32(length) != receivedHeader.Length {
			log.Printf("recieved length(%d) is mismatching with metadata length(%d)\n", length, receivedHeader.Length)
			return
		}

		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("Got", receivedHeader.Type.String())

	for _, listener := range c.listeners[receivedHeader.Type] {
		listener.emit(payload)
	}
}

func (c *Client) StartHandleEvents() {
	go func() {
		for {
			c.HandleEvents()
		}
	}()
}

func Connect(host string, port uint16) (*Client, *version.Version, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		return nil, nil, errors.New("could not connect to server")
	}

	client := &Client{
		conn:      conn,
		listeners: map[events.EventType][]*Listener{},
	}

	client.StartHandleEvents()

	pong, err := client.Send(clientbuffer.PingBuffer()).Wait(events.Pong)

	if err != nil {
		return nil, nil, err
	}

	return client, version.FromBuffer(pong), nil
}

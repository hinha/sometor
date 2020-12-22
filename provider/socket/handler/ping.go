package handler

import (
	"fmt"
	"github.com/hinha/sometor/provider"
	//"golang.org/x/net/websocket"
	"github.com/gorilla/websocket"
)

// We set our Read and Write buffer sizes
var upgrades = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type PingSocket struct {
}

func NewPing() *PingSocket {
	return &PingSocket{}
}

// Path return api path
func (h *PingSocket) Path() string {
	return "/ping"
}

// Method return api method
func (h *PingSocket) Method() string {
	return "GET"
}

// Handle health which always return 200
func (h *PingSocket) Handle(context provider.SocketContext) {
	ws, err := upgrades.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			context.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			context.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

package socket

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/hinha/sometor/provider"
	"log"
	"time"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 10 * time.Second

	// TextMessage denotes a text data message. The text message payload is
	// interpreted as UTF-8 encoded text data.
	TextMessage = 1

	// BinaryMessage denotes a binary data message.
	BinaryMessage = 2

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)

// We set our Read and Write buffer sizes
var upgrades = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type TwitterSocketServe struct {
	socketProvider provider.SocketTwitter
}

func NewTwitterSocket(provider provider.SocketTwitter) *TwitterSocketServe {
	return &TwitterSocketServe{socketProvider: provider}
}

// Path return socket path
func (t *TwitterSocketServe) Path() string {
	return "/twitter"
}

// Method return api method
func (t *TwitterSocketServe) Method() string {
	return "GET"
}

// Handle health which always return 200
func (t *TwitterSocketServe) Handle(context provider.SocketContext) {
	ws, err := upgrades.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	var lastMod time.Time

	for {
		userID := context.QueryParam("id")
		keyword := context.QueryParam("keyword")
		if userID == "" || keyword == "" {
			msg, _ := json.Marshal(map[string]interface{}{
				"errors":  "bad request given by client",
				"message": "Bad request",
			})
			err := ws.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				context.Logger().Error(err)
			}
			ws.Close()
			return
		} else {
			// 1. get valid user id
			// 2. select file by keyword valid
			_, err := t.socketProvider.UserValid(context.Request().Context(), userID, keyword, "twitter")
			if err != nil {
				msg, _ := json.Marshal(map[string]interface{}{
					"errors":  err.Error(),
					"message": "Bad request",
				})
				err := ws.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					context.Logger().Error(err)
				}
				ws.Close()
				return
			} else {
				go t.writer(context.Request().Context(), ws, lastMod, keyword, "twitter")
				ws.PingHandler()
				err := t.reader(ws)
				if err != nil {
					break
				}
			}
		}

	}
}

func (t *TwitterSocketServe) writer(ctx context.Context, ws *websocket.Conn, lastMod time.Time, keyword, media string) {
	lastError := ""
	pingTicker := time.NewTicker(pingPeriod)
	fileTicker := time.NewTicker(filePeriod)

	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()

	for {
		select {
		case <-fileTicker.C:
			var p []byte
			var err error

			p, lastMod, err = t.socketProvider.FileReader(ctx, lastMod, media, keyword)
			if err != nil {
				if s := err.Error(); s != lastError {
					lastError = s
					p = []byte(lastError)
				}
			} else {
				lastError = ""
			}

			if p != nil {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.WriteMessage(TextMessage, p); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(PingMessage, []byte{}); err != nil {
				return
			}
		}

	}
}

func (t *TwitterSocketServe) reader(ws *websocket.Conn) error {
	defer func() {
		ws.Close()
	}()

	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			return err
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			return err
		}
		//fmt.Println("reader: ", ws.Close())
		//if err := ws.Close(); err != nil {
		//	break
		//}
	}
}

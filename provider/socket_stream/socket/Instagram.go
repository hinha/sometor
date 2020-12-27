package socket

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"log"
	"time"
)

type InstagramSocketServe struct {
	socketProvider provider.SocketMedia
}

func NewInstagramSocket(provider provider.SocketMedia) *InstagramSocketServe {
	return &InstagramSocketServe{socketProvider: provider}
}

// Path return socket path
func (t *InstagramSocketServe) Path() string {
	return "/instagram"
}

// Method return api method
func (t *InstagramSocketServe) Method() string {
	return "GET"
}

// Handle health which always return 200
func (t *InstagramSocketServe) Handle(context provider.SocketContext) {
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
			_, err := t.socketProvider.UserValid(context.Request().Context(), userID, keyword, "instagram")
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
				go t.writer(context.Request().Context(), ws, lastMod, keyword, "instagram")
				ws.PingHandler()
				err := t.reader(ws)
				if err != nil {
					break
				}
			}
		}

	}
}

func (t *InstagramSocketServe) writer(ctx context.Context, ws *websocket.Conn, lastMod time.Time, keyword, media string) {
	//lastError := ""
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
			var err *entity.ApplicationError

			p, lastMod, err = t.socketProvider.FileReader(ctx, lastMod, media, keyword)
			if err != nil {
				msg, _ := json.Marshal(map[string]interface{}{
					"message": err.Error(),
				})
				p = msg
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

func (t *InstagramSocketServe) reader(ws *websocket.Conn) error {
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
	}
}

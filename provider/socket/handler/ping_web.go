package handler

import (
	"github.com/hinha/sometor/provider"
	"net/http"
)

type PingWeb struct {
}

func NewPingWeb() *PingWeb {
	return &PingWeb{}
}

// Path return api path
func (h *PingWeb) Path() string {
	return "/test"
}

// Method return api method
func (h *PingWeb) Method() string {
	return "GET"
}

// Handle health which always return 200
func (h *PingWeb) Handle(context provider.SocketContext) {
	_ = context.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

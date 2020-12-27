package socket

import (
	"github.com/hinha/sometor/provider"
	"net/http"
)

type TwitterWebDemo struct {
	socketProvider provider.SocketMedia
}

func NewTwitterDemoWeb(provider provider.SocketMedia) *TwitterWebDemo {
	return &TwitterWebDemo{socketProvider: provider}
}

// Path return api path
func (h *TwitterWebDemo) Path() string {
	return "/demo/twitter"
}

// Method return api method
func (h *TwitterWebDemo) Method() string {
	return "GET"
}

// Handle health which always return 200
func (h *TwitterWebDemo) Handle(context provider.SocketContext) {
	_ = context.Render(http.StatusOK, "index_twitter.html", map[string]interface{}{
		"title": "Test Twitter",
	})
}

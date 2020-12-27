package socket

import (
	"github.com/hinha/sometor/provider"
	"net/http"
)

type InstagramWebDemo struct {
	socketProvider provider.SocketMedia
}

func NewInstagramDemoWeb(provider provider.SocketMedia) *InstagramWebDemo {
	return &InstagramWebDemo{socketProvider: provider}
}

// Path return api path
func (h *InstagramWebDemo) Path() string {
	return "/demo/instagram"
}

// Method return api method
func (h *InstagramWebDemo) Method() string {
	return "GET"
}

// Handle health which always return 200
func (h *InstagramWebDemo) Handle(context provider.SocketContext) {
	_ = context.Render(http.StatusOK, "index_instagram.html", map[string]interface{}{
		"title": "Test Instagram",
	})
}

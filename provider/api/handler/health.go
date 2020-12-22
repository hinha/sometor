package handler

import (
	"github.com/hinha/sometor/provider"
	"net/http"
)

// Health api handler
type Health struct{}

// NewHealth create new health object
func NewHealth() *Health {
	return &Health{}
}

// Path return api path
func (h *Health) Path() string {
	return "/v1/health"
}

// Method return api method
func (h *Health) Method() string {
	return "GET"
}

// Handle health which always return 200
func (h *Health) Handle(context provider.APIContext) {
	_ = context.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

package api

import (
	"github.com/hinha/sometor/provider"
	"net/http"
)

// List keyword api handler
type ListStreamKeyword struct {
	streamProvider provider.SocmedKeywordAPI
}

// NewListStreamKeyword list keyword handler object
func NewListStreamKeyword(streamProvider provider.SocmedKeywordAPI) *ListStreamKeyword {
	return &ListStreamKeyword{streamProvider: streamProvider}
}

// Path return api path
func (l *ListStreamKeyword) Path() string {
	return "/stream/keyword/list"
}

// Method return api method
func (l *ListStreamKeyword) Method() string {
	return "GET"
}

// Handle request list keyword
func (l *ListStreamKeyword) Handle(context provider.APIContext) {
	userID := context.QueryParam("user_id")
	if userID == "" {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}
	response, err := l.streamProvider.StreamKeywordList(context.Request().Context(), userID)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{"data": response})
}

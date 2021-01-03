package api

import (
	"encoding/json"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

// List keyword api handler
type CreateStreamKeyword struct {
	streamProvider provider.SocmedKeywordAPI
}

// NewCreateStreamKeyword list keyword handler object
func NewCreateStreamKeyword(streamProvider provider.SocmedKeywordAPI) *CreateStreamKeyword {
	return &CreateStreamKeyword{streamProvider: streamProvider}
}

// Path return api path
func (l *CreateStreamKeyword) Path() string {
	return "/stream/keyword/create"
}

// Method return api method
func (l *CreateStreamKeyword) Method() string {
	return "POST"
}

// Handle request create keyword
func (l *CreateStreamKeyword) Handle(context provider.APIContext) {
	var request entity.StreamSequenceInsertable
	if err := json.NewDecoder(context.Request().Body).Decode(&request); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	if request.Keyword == "" || request.Media == "" || request.Type == "" || request.UserAccountID == "" {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	if len(request.Keyword) > 60 || len(request.UserAccountID) > 50 {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	switch request.Media {
	case "twitter", "instagram":
		response, err := l.streamProvider.StreamKeywordCreate(context.Request().Context(), request)
		if err != nil {
			_ = context.JSON(err.HTTPStatus, map[string]interface{}{
				"errors":  err.ErrorString(),
				"message": err.Error(),
			})
			return
		}

		_ = context.JSON(http.StatusOK, map[string]interface{}{"data": response})
		return
	default:
		_ = context.JSON(http.StatusNotFound, map[string]interface{}{
			"errors":  []string{"request media not found"},
			"message": "NotFound",
		})
		return
	}

}

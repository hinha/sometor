package api

import (
	"fmt"
	"github.com/hinha/sometor/provider"
	"net/http"
	"strconv"
)

// Delete keyword api handler
type DeleteStreamKeyword struct {
	streamProvider provider.SocmedKeywordAPI
}

// NewDeleteStreamKeyword delete keyword handler object
func NewDeleteStreamKeyword(streamProvider provider.SocmedKeywordAPI) *DeleteStreamKeyword {
	return &DeleteStreamKeyword{streamProvider: streamProvider}
}

// Path return api path
func (l *DeleteStreamKeyword) Path() string {
	return "/stream/keyword/delete"
}

// Method return api method
func (l *DeleteStreamKeyword) Method() string {
	return "DELETE"
}

// Handle request delete keyword by id
func (l *DeleteStreamKeyword) Handle(context provider.APIContext) {
	id := context.QueryParam("keyID")
	userID := context.QueryParam("user_id")
	fmt.Println(id, userID)
	if id == "" || userID == "" {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	if len(id) > 30 || len(userID) > 50 {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	keyID, err := strconv.Atoi(id)
	if err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request id not supported string"},
			"message": "Bad request",
		})
		return
	}

	result, errProvider := l.streamProvider.StreamKeywordDelete(context.Request().Context(), keyID, userID)
	if errProvider != nil {
		_ = context.JSON(errProvider.HTTPStatus, map[string]interface{}{
			"errors":  errProvider.ErrorString(),
			"message": errProvider.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"status": result,
		},
		"message": "successfully delete",
	})

	return
}

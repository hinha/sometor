package api

import (
	"fmt"
	"github.com/hinha/sometor/provider"
	"net/http"
)

// Show keyword data api handler
type UpdateStreamData struct {
	streamProvider provider.SocmedKeywordAPI
}

// NewUpdateStreamData show data stream keyword handler object
func NewUpdateStreamData(streamProvider provider.SocmedKeywordAPI) *UpdateStreamData {
	return &UpdateStreamData{streamProvider: streamProvider}
}

// Path return api path
func (s *UpdateStreamData) Path() string {
	return "/stream/:media/update"
}

// Method return api method
func (s *UpdateStreamData) Method() string {
	return "GET"
}

// Handle request create keyword
func (s *UpdateStreamData) Handle(context provider.APIContext) {
	userID := context.QueryParam("user_id")
	userKeyword := context.QueryParam("keyword")

	if userID == "" || userKeyword == "" {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	if len(userID) > 100 || len(userKeyword) > 120 {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad length request given by client"},
			"message": "Bad request",
		})
		return
	}

	mediaType := context.Param("media")
	switch mediaType {
	case "twitter":
		data, errProvider := s.streamProvider.StreamKeywordUpdateDataTwitter(context.Request().Context(), mediaType, userID, userKeyword)
		if errProvider != nil {
			_ = context.JSON(errProvider.HTTPStatus, map[string]interface{}{
				"errors":  errProvider.ErrorString(),
				"message": errProvider.Error(),
			})
			return
		}

		_ = context.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
		return
	case "instagram":
		data, errProvider := s.streamProvider.StreamKeywordUpdateDataInstagram(context.Request().Context(), mediaType, userID, userKeyword)
		if errProvider != nil {
			_ = context.JSON(errProvider.HTTPStatus, map[string]interface{}{
				"errors":  errProvider.ErrorString(),
				"message": errProvider.Error(),
			})
			return
		}

		_ = context.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
		return
	default:
		_ = context.JSON(http.StatusNotFound, map[string]interface{}{
			"errors":  []string{fmt.Sprintf("no media found %s", context.Param("media"))},
			"message": "Bad request",
		})
		return
	}
}

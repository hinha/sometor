package api

import (
	"fmt"
	"github.com/hinha/sometor/provider"
	"net/http"
)

// Show keyword data api handler
type ShowStreamData struct {
	streamProvider provider.SocmedKeywordAPI
}

// NewShowStreamData show data stream keyword handler object
func NewShowStreamData(streamProvider provider.SocmedKeywordAPI) *ShowStreamData {
	return &ShowStreamData{streamProvider: streamProvider}
}

// Path return api path
func (s *ShowStreamData) Path() string {
	return "/stream/:media/show"
}

// Method return api method
func (s *ShowStreamData) Method() string {
	return "GET"
}

// Handle request create keyword
func (s *ShowStreamData) Handle(context provider.APIContext) {
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
		data, errProvider := s.streamProvider.StreamKeywordShowDataTwitter(context.Request().Context(), mediaType, userID, userKeyword)
		if errProvider != nil {
			if errProvider.HTTPStatus == http.StatusOK {
				_ = context.JSON(http.StatusOK, map[string]interface{}{
					"data": data,
				})
				return
			}

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
		data, errProvider := s.streamProvider.StreamKeywordShowDataInstagram(context.Request().Context(), mediaType, userID, userKeyword)
		if errProvider != nil {
			if errProvider.HTTPStatus == http.StatusOK {
				_ = context.JSON(http.StatusOK, map[string]interface{}{
					"data": data,
				})
				return
			}
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
	case "facebook":
		data, errProvider := s.streamProvider.StreamKeywordShowDataFacebook(context.Request().Context(), mediaType, userID, userKeyword)
		if errProvider != nil {
			if errProvider.HTTPStatus == http.StatusOK {
				_ = context.JSON(http.StatusOK, map[string]interface{}{
					"data": data,
				})
				return
			}
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

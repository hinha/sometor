package api

import (
	"github.com/hinha/sometor/provider"
	"net/http"
)

type listStatusesTwitterOauth struct {
	streamProvider provider.SocmedKeywordAPI
}

func NewListStatusesTwitterOauth(streamProvider provider.SocmedKeywordAPI) *listStatusesTwitterOauth {
	return &listStatusesTwitterOauth{streamProvider: streamProvider}
}

// Path return api path
func (l *listStatusesTwitterOauth) Path() string {
	return "/authorization/twitter/statuses"
}

// Method return api method
func (l *listStatusesTwitterOauth) Method() string {
	return "GET"
}

func (l *listStatusesTwitterOauth) Handle(context provider.APIContext) {
	userTweetID := context.QueryParam("user_tweet_id")
	userID := context.QueryParam("user_id")

	if userTweetID == "" || userID == "" {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	statuses, errProvider := l.streamProvider.TwitterListStatuses(context.Request().Context(), userTweetID, userID)
	if errProvider != nil {
		_ = context.JSON(errProvider.HTTPStatus, map[string]interface{}{
			"errors":  errProvider.ErrorString(),
			"message": errProvider.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{
		"data": statuses,
	})
	return
}

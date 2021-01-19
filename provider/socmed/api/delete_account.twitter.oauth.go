package api

import (
	"github.com/hinha/sometor/provider"
	"net/http"
)

type delAccountTwitterOauth struct {
	streamProvider provider.SocmedKeywordAPI
}

func NewDelAccountTwitterOauth(streamProvider provider.SocmedKeywordAPI) *delAccountTwitterOauth {
	return &delAccountTwitterOauth{streamProvider: streamProvider}
}

// Path return api path
func (d *delAccountTwitterOauth) Path() string {
	return "/authorization/twitter/user"
}

// Method return api method
func (d *delAccountTwitterOauth) Method() string {
	return "DELETE"
}

func (d *delAccountTwitterOauth) Handle(context provider.APIContext) {
	userTweetID := context.QueryParam("user_tweet_id")
	userID := context.QueryParam("user_id")

	if userTweetID == "" || userID == "" {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	result, errProvider := d.streamProvider.TwitterDelAccountOauth(context.Request().Context(), userTweetID, userID)
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

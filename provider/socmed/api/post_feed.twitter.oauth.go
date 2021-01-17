package api

import (
	"encoding/json"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type postFeedTwitterOauth struct {
	streamProvider provider.SocmedKeywordAPI
}

func NewPostFeedTwitter(streamProvider provider.SocmedKeywordAPI) *postFeedTwitterOauth {
	return &postFeedTwitterOauth{streamProvider: streamProvider}
}

// Path return api path
func (p *postFeedTwitterOauth) Path() string {
	return "/authorization/twitter/statuses"
}

// Method return api method
func (p *postFeedTwitterOauth) Method() string {
	return "POST"
}

func (p *postFeedTwitterOauth) Handle(context provider.APIContext) {
	var bodyJson entity.OFeedTwitter

	if err := json.NewDecoder(context.Request().Body).Decode(&bodyJson); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	err := p.streamProvider.TwitterPostFeedOauth(context.Request().Context(), bodyJson)
	if err != nil {
		_ = context.JSON(err.HTTPStatus, map[string]interface{}{
			"errors":  err.ErrorString(),
			"message": err.Error(),
		})
		return
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{
		"message": "tweet has been posted",
		"status":  "ok",
	})

}

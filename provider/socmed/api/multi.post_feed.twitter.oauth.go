package api

import (
	"encoding/json"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type postFeedTwitterOauthMulti struct {
	streamProvider provider.SocmedKeywordAPI
}

func NewPostFeedTwitterMulti(streamProvider provider.SocmedKeywordAPI) *postFeedTwitterOauthMulti {
	return &postFeedTwitterOauthMulti{streamProvider: streamProvider}
}

// Path return api path
func (p *postFeedTwitterOauthMulti) Path() string {
	return "/authorization/twitter/m/statuses"
}

// Method return api method
func (p *postFeedTwitterOauthMulti) Method() string {
	return "POST"
}

func (p *postFeedTwitterOauthMulti) Handle(context provider.APIContext) {
	var bodyJson entity.OFeedTwitterAll

	if err := json.NewDecoder(context.Request().Body).Decode(&bodyJson); err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	err := p.streamProvider.TwitterPostFeedAllOauth(context.Request().Context(), bodyJson)
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

package api

import (
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type listAccountTwitterOauth struct {
	streamProvider provider.SocmedKeywordAPI
}

func NewListAccountTwitterOauth(streamProvider provider.SocmedKeywordAPI) *listAccountTwitterOauth {
	return &listAccountTwitterOauth{streamProvider: streamProvider}
}

// Path return api path
func (l *listAccountTwitterOauth) Path() string {
	return "/authorization/twitter/list"
}

// Method return api method
func (l *listAccountTwitterOauth) Method() string {
	return "GET"
}

func (l *listAccountTwitterOauth) Handle(context provider.APIContext) {
	userID := context.QueryParam("user_id")
	if userID == "" {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	users, errProvider := l.streamProvider.TwitterListOauthAccount(context.Request().Context(), userID)
	if errProvider != nil {
		_ = context.JSON(errProvider.HTTPStatus, map[string]interface{}{
			"errors":  errProvider.ErrorString(),
			"message": errProvider.Error(),
		})
		return
	}

	if len(users) == 0 {
		users = []entity.OUserTwitterInfo{}
	}

	_ = context.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
	return
}

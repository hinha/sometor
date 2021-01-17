package api

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
	"regexp"
)

// List keyword api handler
type OauthCallback struct {
	streamProvider provider.SocmedKeywordAPI
}

// NewOauthCallback list keyword handler object
func NewOauthCallback(streamProvider provider.SocmedKeywordAPI) *OauthCallback {
	return &OauthCallback{streamProvider: streamProvider}
}

// Path return api path
func (o *OauthCallback) Path() string {
	return "/authorization/:provider/client"
}

// Method return api method
func (o *OauthCallback) Method() string {
	return "GET"
}

//ewogICAgIm5hbWUiOiAiTWFydGluIiwKICAgICJ1c2VybmFtZSI6ICJtYXJsdGlubCIsCiAgICAicHJvZmlsZV9pbWFnZV91cmwiOiAiaHR0cHM6Ly9wYnMudHdpbWcuY29tL3Byb2ZpbGVfaW1hZ2VzLzEyNDY0OTY2NDQxNjc2NTU0MjUvWWpsd1RyYlRfbm9ybWFsLmpwZywKICAgICJ1c2VyX2lkIjogIjExMDY4MTc3Mzg2ODM1Mzk1MDAiLAogICAgImFjY2Vzc190b2tlbiI6ICIxMTA2ODE3NzM4NjgzNTM5NDU2LWN0Y0podUU0ZDF2Y3ZnVnVuMGpkeXlSN3lzTnJmUiIsCiAgICAiYWNjZXNzX3Rva2VuX3NlY3JldCI6ICJpNjcwR0dJemlwakU4aFNyaWpCaDJ3YTZMRDh5cFJEUkFtWFRESFlTQ3AxM2MiLAogICAgImNsaWVudF9pZCI6ICJlMWRmNmQ0YjUzMjBkMjRjNTQwMGE3NTZhODFiNzJhYThlMGFhNjQ1OGI0YTE2Igp9
// Handle request list keyword
func (o *OauthCallback) Handle(context provider.APIContext) {
	var uTwitter entity.OUserTwitter
	payLoad := context.QueryParam("payload")

	if payLoad == "" {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	pload64, err := b64.StdEncoding.DecodeString(payLoad)
	if err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"payload error"},
			"message": "Bad request",
		})
		return
	}

	space := regexp.MustCompile(`\s+`)
	text := space.ReplaceAllString(string(pload64), "")
	err = json.Unmarshal([]byte(text), &uTwitter)
	if err != nil {
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{"bad request given by client"},
			"message": "Bad request",
		})
		return
	}

	switch context.Param("provider") {
	case "twitter":
		user, errProvider := o.streamProvider.TwitterOauthToken(context.Request().Context(), uTwitter)
		if errProvider != nil {
			_ = context.JSON(errProvider.HTTPStatus, map[string]interface{}{
				"errors":  errProvider.ErrorString(),
				"message": errProvider.Error(),
			})
			return
		}
		_ = context.JSON(http.StatusOK, map[string]interface{}{
			"data": user,
		})
		return
	default:
		_ = context.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors":  []string{fmt.Sprintf("provider not found %s", context.Param("provider"))},
			"message": "Bad request",
		})
		return
	}
}

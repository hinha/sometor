package api

import (
	"fmt"
	"github.com/hinha/sometor/provider"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitter"
	"net/http"
	"sort"
)

// List keyword api handler
type OauthClientToken struct {
	streamProvider provider.SocmedKeywordAPI
}

// NewOauthClientToken list keyword handler object
func NewOauthClientToken(streamProvider provider.SocmedKeywordAPI) *OauthClientToken {
	return &OauthClientToken{streamProvider: streamProvider}
}

// Path return api path
func (o *OauthClientToken) Path() string {
	return "/authorization/:provider/client_token"
}

// Method return api method
func (o *OauthClientToken) Method() string {
	return "GET"
}

func (o *OauthClientToken) Handle(context provider.APIContext) {

	goth.UseProviders(
		twitter.New("ROaXlUvhpKXxTDVTyXH3tKyOk", "guuG1z0QLtUG6Aea28Hlub0hOgW5Ps0jh66u1YdSNhbUvyAP16", "http://localhost:9091/auth/twitter/callback"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		//twitter.NewAuthenticate("ROaXlUvhpKXxTDVTyXH3tKyOk", "guuG1z0QLtUG6Aea28Hlub0hOgW5Ps0jh66u1YdSNhbUvyAP16", "http://localhost:3000/auth/twitter/callback"),
	)

	m := make(map[string]string)
	m["twitter"] = "Twitter"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	switch context.Param("provider") {
	case "twitter":
		//user, errProvider := o.streamProvider.TwitterOauthToken(context.Request().Context(), uTwitter)
		//if errProvider != nil {
		//	_ = context.JSON(errProvider.HTTPStatus, map[string]interface{}{
		//		"errors":  errProvider.ErrorString(),
		//		"message": errProvider.Error(),
		//	})
		//	return
		//}
		res := context.Response()
		user, err := gothic.CompleteUserAuth(res.Writer, context.Request())
		if err != nil {
			fmt.Fprintln(res.Writer, err)
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

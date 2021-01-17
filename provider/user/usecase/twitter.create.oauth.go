package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type CreateOauthTw struct{}

func (c *CreateOauthTw) Perform(ctx context.Context, request entity.OUserTwitter, db provider.DB) *entity.ApplicationError {
	result, err := db.ExecContext(ctx, "create-oauth-tw", "INSERT INTO oauth_twitter (user_id,name,username,profile_image_url,access_token,access_token_secret,user_account_id,created_at) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, now());",
		request.UserID, request.Name, request.Username, request.ProfileImageURL, request.AccessToken, request.AccessTokenSecret, request.ClientID)

	if err != nil {
		return &entity.ApplicationError{
			Err:        []error{errors.New("user not found")},
			HTTPStatus: http.StatusBadRequest,
		}
	}

	_, err = result.LastInsertId()
	if err != nil {
		return &entity.ApplicationError{
			Err:        []error{errors.New("internal server error when acquiring last inserted id")},
			HTTPStatus: http.StatusInternalServerError,
		}
	}

	return nil
}

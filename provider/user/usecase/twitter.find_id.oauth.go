package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type FindByIDOauthTw struct{}

func (u *FindByIDOauthTw) Perform(ctx context.Context, ID string, db provider.DB) (entity.OUserTwitterInfo, *entity.ApplicationError) {
	var data entity.OUserTwitterInfo
	row := db.QueryRowContext(ctx, "find-stream-user_account", "select user_id,name,username,profile_image_url from oauth_twitter where user_id = ?", ID)
	if err := row.Scan(&data.UserTweetID, &data.Name, &data.Username, &data.ProfileImageURL); err == provider.ErrDBNotFound {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("user not found")},
			HTTPStatus: http.StatusNotFound,
		}
	} else if err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	return data, nil
}

type FindByIdKeyOauthTw struct{}

func (f *FindByIdKeyOauthTw) Perform(ctx context.Context, IdTweet, UserID string, db provider.DB) (entity.OUserTwitterKey, *entity.ApplicationError) {
	var data entity.OUserTwitterKey
	row := db.QueryRowContext(ctx, "find-stream-user_account", "select user_id,access_token,access_token_secret from oauth_twitter where user_id = ? and user_account_id = ?", IdTweet, UserID)
	if err := row.Scan(&data.UserTweetID, &data.AccessToken, &data.AccessTokenSecret); err == provider.ErrDBNotFound {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("user not found")},
			HTTPStatus: http.StatusNotFound,
		}
	} else if err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	return data, nil
}

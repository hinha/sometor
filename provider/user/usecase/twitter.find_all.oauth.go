package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type FindAllOauthTw struct{}

func (f *FindAllOauthTw) Perform(ctx context.Context, userID string, db provider.DB) ([]entity.OUserTwitterInfo, *entity.ApplicationError) {
	var data []entity.OUserTwitterInfo
	rows, err := db.QueryContext(ctx, "find-stream-twitt_acc", "select user_id,name,username,profile_image_url from oauth_twitter where user_account_id = ?", userID)
	if err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("twitter account not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	for rows.Next() {
		var each entity.OUserTwitterInfo
		if err := rows.Scan(&each.UserTweetID, &each.Name, &each.Username, &each.ProfileImageURL); err != nil {
			return data, &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			}
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("[rows] service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	return data, nil
}

type FindAllOauthTwKey struct{}

func (f *FindAllOauthTwKey) Perform(ctx context.Context, userID string, db provider.DB) ([]entity.OUserTwitterKey, *entity.ApplicationError) {
	var data []entity.OUserTwitterKey
	rows, err := db.QueryContext(ctx, "find-stream-tweet_key", "select user_id,access_token,access_token_secret from oauth_twitter where user_account_id = ?", userID)
	if err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("twitter account not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	for rows.Next() {
		var each entity.OUserTwitterKey
		if err := rows.Scan(&each.UserTweetID, &each.AccessToken, &each.AccessTokenSecret); err != nil {
			return data, &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			}
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("[rows] service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
		}
	}

	return data, nil
}

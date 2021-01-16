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
	if err := row.Scan(&data.UserID, &data.Name, &data.Username, &data.ProfileImageURL); err == provider.ErrDBNotFound {
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

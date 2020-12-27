package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type StreamRequestByID struct{}

func (s *StreamRequestByID) Perform(ctx context.Context, ID string, db provider.DB) (entity.StreamSequenceInitTable, *entity.ApplicationError) {
	var data entity.StreamSequenceInitTable
	row := db.QueryRowContext(ctx, "find-stream-user", "select id,keyword,media,type,created_at,user_account_id from stream_sequence_account where user_account_id = ?", ID)
	if err := row.Scan(
		&data.ID,
		&data.Keyword,
		&data.Media,
		&data.Type,
		&data.CreatedAt,
		&data.UserAccountID); err == provider.ErrDBNotFound {
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

type UserAccountID struct{}

func (u *UserAccountID) PerformInfo(ctx context.Context, ID string, db provider.DB) (entity.UserAccountSelectable, *entity.ApplicationError) {
	var data entity.UserAccountSelectable
	row := db.QueryRowContext(ctx, "find-stream-user_account", "select unique_account,email,unique_user from user_account where unique_account = ?", ID)
	if err := row.Scan(&data.UniqueAccount, &data.Email, &data.UniqueUser); err == provider.ErrDBNotFound {
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

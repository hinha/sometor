package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type FindByKeywordStream struct{}

func (f *FindByKeywordStream) Perform(ctx context.Context, fileName, media string, db provider.DB) (entity.StreamSequenceInitTable, *entity.ApplicationError) {
	var data entity.StreamSequenceInitTable
	row := db.QueryRowContext(ctx, "find-stream-user", "select id,keyword,media,type,created_at,user_account_id from stream_sequence_account where keyword = ? and media = ?", fileName, media)
	if err := row.Scan(
		&data.ID,
		&data.Keyword,
		&data.Media,
		&data.Type,
		&data.CreatedAt,
		&data.UserAccountID); err == provider.ErrDBNotFound {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("keyword not found")},
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

func (f *FindByKeywordStream) PerformWithAccount(ctx context.Context, keyword, types, media, UserID string, db provider.DB) (entity.StreamSequenceInsertable, *entity.ApplicationError) {
	var data entity.StreamSequenceInsertable
	row := db.QueryRowContext(ctx, "find-stream-user-l", "select keyword,type,user_account_id,media from stream_sequence_account where keyword = ? and type = ? and user_account_id = ? and media = ?", keyword, media, UserID, types)
	if err := row.Scan(
		&data.Keyword, &data.Type, &data.UserAccountID, &data.Media); err == provider.ErrDBNotFound {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("keyword not found")},
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

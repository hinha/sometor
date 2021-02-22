package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type StreamRequestAll struct{}

func (s *StreamRequestAll) Perform(ctx context.Context, db provider.DB) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	var data []entity.StreamSequenceInitTable
	rows, err := db.QueryContext(ctx, "find-stream-user", "select id,keyword,media,type,created_at,user_account_id from stream_sequence_account group by keyword order by created_at DESC")
	if err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("account not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	for rows.Next() {
		var each entity.StreamSequenceInitTable
		if err := rows.Scan(&each.ID,
			&each.Keyword,
			&each.Media,
			&each.Type,
			&each.CreatedAt,
			&each.UserAccountID); err != nil {

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

type StreamRequestMedia struct{}

func (s *StreamRequestMedia) Perform(ctx context.Context, media string, db provider.DB) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	var data []entity.StreamSequenceInitTable
	rows, err := db.QueryContext(ctx, "find-stream-user", "select id,keyword,media,type,created_at,user_account_id from stream_sequence_account where media = ? "+
		"group by keyword order by created_at DESC", media)
	if err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("account not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	for rows.Next() {
		var each entity.StreamSequenceInitTable
		if err := rows.Scan(&each.ID,
			&each.Keyword,
			&each.Media,
			&each.Type,
			&each.CreatedAt,
			&each.UserAccountID); err != nil {

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

type StreamRequestID struct{}

func (s *StreamRequestID) Perform(ctx context.Context, ID string, db provider.DB) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	var data []entity.StreamSequenceInitTable
	rows, err := db.QueryContext(ctx, "find-stream-user", "select id,keyword,media,type,created_at,user_account_id from stream_sequence_account where user_account_id = ? order by created_at DESC", ID)
	if err != nil {
		return data, &entity.ApplicationError{
			Err:        []error{errors.New("account not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	for rows.Next() {
		var each entity.StreamSequenceInitTable
		if err := rows.Scan(&each.ID,
			&each.Keyword,
			&each.Media,
			&each.Type,
			&each.CreatedAt,
			&each.UserAccountID); err != nil {

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

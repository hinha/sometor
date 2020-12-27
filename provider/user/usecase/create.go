package usecase

import (
	"context"
	"errors"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type CreateKeywordStream struct{}

func (c *CreateKeywordStream) Perform(ctx context.Context, request entity.StreamSequenceInsertable, db provider.DB) (int, *entity.ApplicationError) {
	result, err := db.ExecContext(ctx, "create-keyword-stream", "INSERT INTO stream_sequence_account (keyword,media,type,created_at,user_account_id) VALUES (?,?,?,now(),?)",
		request.Keyword, request.Media, request.Type, request.UserAccountID)

	if err != nil {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("user not found")},
			HTTPStatus: http.StatusBadRequest,
		}
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("internal server error when acquiring last inserted id")},
			HTTPStatus: http.StatusInternalServerError,
		}
	}

	return int(ID), nil
}

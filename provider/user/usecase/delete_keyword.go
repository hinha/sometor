package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type DeleteKeywordStream struct{}

func (d *DeleteKeywordStream) Perform(ctx context.Context, ID int, userID string, db provider.DB) (int, *entity.ApplicationError) {
	format := fmt.Sprintf("DELETE FROM stream_sequence_account WHERE id = %d and user_account_id = '%s'", ID, userID)
	result, err := db.ExecContext(ctx, "delete-keyword-stream", format)
	if err != nil {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("keyword not found")},
			HTTPStatus: http.StatusBadRequest,
		}
	}

	IDn, err := result.RowsAffected()
	if err != nil {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("internal server error when acquiring id")},
			HTTPStatus: http.StatusInternalServerError,
		}
	}
	return int(IDn), nil
}

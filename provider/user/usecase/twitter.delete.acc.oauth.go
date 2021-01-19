package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type TwitterDelAccountOauth struct{}

func (t *TwitterDelAccountOauth) Perform(ctx context.Context, tweetUserID, userID string, db provider.DB) (int, *entity.ApplicationError) {
	format := fmt.Sprintf("DELETE FROM oauth_twitter WHERE user_id = '%s' and user_account_id = '%s'", tweetUserID, userID)
	result, err := db.ExecContext(ctx, "delete-tweet-account", format)
	if err != nil {
		return 0, &entity.ApplicationError{
			Err:        []error{errors.New("tweet account not found")},
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

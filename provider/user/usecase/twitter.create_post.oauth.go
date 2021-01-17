package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"net/http"
)

type TwitterCreatePostOauth struct{}

func (t *TwitterCreatePostOauth) Perform(ctx context.Context, request entity.OFeedTwitter, statusID, username string, db provider.DB) *entity.ApplicationError {
	permalink := fmt.Sprintf("https://twitter.com/%s/status/%s", username, statusID)
	result, err := db.ExecContext(ctx, "create-keyword-stream", "INSERT INTO twitter_statuses (statuses_id,text,lang,permalink,created_at,user_id) VALUES (?, ?, ?, ?, now(), ?);",
		statusID, request.Text, "id", permalink, request.UserTweetID)

	if err != nil {
		return &entity.ApplicationError{
			Err:        []error{errors.New("service unavailable")},
			HTTPStatus: http.StatusServiceUnavailable,
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
